package main

import (
	"context"
)

// App struct
type App struct {
	ctx       context.Context
	simulator *Simulator
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		simulator: NewSimulator(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.simulator.InitializeScene()
	a.simulator.Start(ctx)
}

// GetSimulationState returns the current simulation state
func (a *App) GetSimulationState() SimulationState {
	return a.simulator.GetState()
}

// SetSimulationSpeed sets the speed multiplier
func (a *App) SetSimulationSpeed(multiplier float64) {
	a.simulator.SetSpeed(multiplier)
}

// GetSimulationSpeed returns the current speed multiplier
func (a *App) GetSimulationSpeed() float64 {
	return a.simulator.GetSpeed()
}
