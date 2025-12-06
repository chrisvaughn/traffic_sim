# Traffic Simulator

A general traffic simulator built with Go and Wails, featuring a Vue/TypeScript frontend. This project explores the complexity and nuance involved in simulating traffic systems, including vehicle dynamics, driver behavior, intersections, and emergent traffic patterns.

## Purpose

This simulator is designed to explore what kind of thinking and complexity goes into modeling traffic systems. Traffic involves many interacting components:
- Vehicle physics (acceleration, braking, speed)
- Driver behavior (reaction times, following distance, decision-making)
- Road networks and intersections
- Traffic control (stop signs, signals)
- Emergent phenomena (jams, shockwaves, flow patterns)

The goal is to build up complexity incrementally to understand how these elements interact and create realistic traffic behavior.

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) (1.18 or later)
- [Node.js](https://nodejs.org/) (16 or later)
- [Wails](https://wails.io/) CLI

Install Wails:
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Development Mode

Run the simulator in live development mode with hot reload:
```bash
wails dev
```

The application will open as a native window. You can also access it in a browser at http://localhost:34115 to debug with devtools.

### Controls
- **Speed buttons (1x, 2x, 5x, 10x)**: Control simulation speed
- The simulation runs continuously, with configurable speed multipliers

## Building

Build a production-ready executable:
```bash
wails build
```

The built application will be in `build/bin/`.

### Project Configuration

Edit `wails.json` to configure project settings. See the [Wails documentation](https://wails.io/docs/reference/project-config) for details.
