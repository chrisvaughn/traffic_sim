package main

import "math"

const (
	HeadingEast  float64 = 0
	HeadingNorth float64 = math.Pi / 2
	HeadingWest  float64 = math.Pi
	HeadingSouth float64 = 3 * math.Pi / 2
)

type Vehicle struct {
	ID       string   `json:"id"`
	Road     *Road    `json:"road"`
	Position Position `json:"position"`
	Velocity float64  `json:"velocity"` // units per second
	Heading  float64  `json:"heading"`  // radians, 0 = East, π/2 = North, π = West, 3π/2 = South
}

func NewVehicle(name string, start Position, velocity, heading float64) Vehicle {
	return Vehicle{
		ID:       name,
		Position: start,
		Velocity: velocity,
		Heading:  heading,
	}
}

func (v *Vehicle) Update(timeDeltaSeconds float64) {
	// Calculate distance traveled this frame
	distance := v.Velocity * timeDeltaSeconds

	// Update position based on heading and distance
	// Heading uses standard math coordinates: 0 = East, π/2 = North
	v.Position.X += distance * math.Cos(v.Heading)
	v.Position.Y -= distance * math.Sin(v.Heading) // Subtract because Y increases downward in canvas
}
