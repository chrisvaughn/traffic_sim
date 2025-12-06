package main

import (
	"context"
	"sync"
	"time"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type StopSign struct {
	ID       string   `json:"id"`
	Position Position `json:"position"`
	RoadID   string   `json:"roadId"`
}

type Intersection struct {
	ID        string     `json:"id"`
	Position  Position   `json:"position"`
	StopSigns []StopSign `json:"stopSigns"`
}

type SimulationState struct {
	Roads         []Road         `json:"roads"`
	Vehicles      []Vehicle      `json:"vehicles"`
	Intersections []Intersection `json:"intersections"`
	Time          float64        `json:"time"` // simulation time in seconds
}

type Simulator struct {
	mu              sync.RWMutex
	state           SimulationState
	speedMultiplier float64
	running         bool
	ticker          *time.Ticker
	cancel          context.CancelFunc
	timestep        time.Duration // fixed timestep for physics
}

func NewSimulator() *Simulator {
	return &Simulator{
		state: SimulationState{
			Roads:         []Road{},
			Vehicles:      []Vehicle{},
			Intersections: []Intersection{},
			Time:          0,
		},
		speedMultiplier: 1.0,
		timestep:        16 * time.Millisecond, // ~60 FPS
		running:         false,
	}
}

func (s *Simulator) Start(ctx context.Context) {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}

	s.running = true
	simCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.mu.Unlock()

	go s.run(simCtx)
}

func (s *Simulator) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	if s.cancel != nil {
		s.cancel()
	}
	s.running = false
}

func (s *Simulator) run(ctx context.Context) {
	ticker := time.NewTicker(s.timestep)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.update()
		}
	}
}

func (s *Simulator) update() {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Calculate actual time delta based on speed multiplier
	dt := s.timestep.Seconds() * s.speedMultiplier

	// Update simulation time
	s.state.Time += dt

	// Update all vehicles
	for i := range s.state.Vehicles {
		s.state.Vehicles[i].Update(dt)
	}

	// TODO: Add more simulation logic
	// - Check for stop signs
	// - Handle intersection logic
	// - Apply physics/behavior rules
}

func (s *Simulator) GetState() SimulationState {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy of the state
	return s.state
}

func (s *Simulator) SetSpeed(multiplier float64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if multiplier < 0 {
		multiplier = 0
	}
	s.speedMultiplier = multiplier
}

func (s *Simulator) GetSpeed() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.speedMultiplier
}

func (s *Simulator) InitializeScene() {
	s.mu.Lock()
	defer s.mu.Unlock()

	roads := []Road{
		NewStraightRoad("north-south", Position{400, 0}, Position{400, 600}),
		NewStraightRoad("east-west", Position{0, 300}, Position{800, 300}),
	}
	s.state.Roads = roads

	// Create intersection at center with 2 stop signs
	intersection := Intersection{
		ID:       "intersection-1",
		Position: Position{X: 400, Y: 300}, // center of a 800x600 canvas
		StopSigns: []StopSign{
			{
				ID:       "stop-1",
				Position: Position{X: 400, Y: 250}, // north approach
				RoadID:   "north-south",
			},
			{
				ID:       "stop-2",
				Position: Position{X: 400, Y: 350}, // south approach
				RoadID:   "north-south",
			},
		},
	}

	s.state.Intersections = []Intersection{intersection}

	vehicles := []Vehicle{
		{
			ID:       "v-1",
			Road:     &s.state.Roads[1],
			Position: Position{X: 0, Y: 300},
			Velocity: 10,
			Heading:  HeadingEast,
		},
	}
	s.state.Vehicles = vehicles
}
