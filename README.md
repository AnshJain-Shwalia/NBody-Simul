# NBody-Simul: N-Body Gravitational Simulation

A visually appealing and scientifically accurate N-Body gravitational simulation written in Go using the Fyne GUI toolkit.

![N-Body Simulation Demo](https://github.com/AnshJain-Shwalia/NBody-Simul/blob/main/demo.gif)

## Overview

This application simulates the gravitational interactions between multiple celestial bodies in a 2D space. Each body exerts gravitational forces on all others according to Newton's law of universal gravitation, resulting in complex orbital patterns and dynamic system behavior.

## Features

- Real-time simulation of gravitational interactions between multiple bodies
- Configurable physics parameters (gravity constant, time step, etc.)
- Random generation of celestial bodies with varying masses
- Visual representation with appropriate scaling
- Concurrent computation for performance optimization

## Requirements

- Go 1.16 or higher
- Fyne v2 (GUI toolkit)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/AnshJain-Shwalia/NBody-Simul.git
   cd NBody-Simul
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the simulation:
   ```
   go run .
   ```

## How It Works

The simulation follows these key steps:

1. **Initialization**: The system creates a set of celestial bodies with random positions, velocities, and masses (configurable via constants).

2. **Force Calculation**: For each time step, the program calculates the gravitational force between each pair of bodies using Newton's law of universal gravitation: F = G * (m1 * m2) / r².

3. **Position Updates**: Based on the net force acting on each body, the simulation updates velocity and position vectors.

4. **Visualization**: The GUI renders each body as a circle with size proportional to its mass, animating movement in real-time.

5. **Concurrency**: The simulation uses Go's goroutines and channels to separate physics calculations from rendering, ensuring smooth animation.

## Configuration

Key parameters can be modified in `constants.go`:

| Parameter | Description |
|-----------|-------------|
| `timeStep` | Simulation time increment per iteration |
| `G` | Gravitational constant |
| `numObjects` | Number of celestial bodies |
| `minDistance`/`maxDistance` | Range for initial distance from center |
| `minVelocity`/`maxVelocity` | Range for initial velocity magnitude |
| `minMass`/`maxMass` | Range for body mass |
| `scaling` | Visual scaling factor |
| `minAccDistance` | Minimum distance for acceleration calculation (prevents division by zero) |

## Architecture

The codebase is organized into several key components:

- `main.go`: Entry point and application setup
- `constants.go`: Simulation parameters and body structure definition
- `positionCalculation.go`: Physics engine for gravitational calculations
- `displayHandling.go`: GUI rendering and animation
- `miscellaneous.go`: Utility functions

## Science Behind the Simulation

This simulation implements Newton's law of universal gravitation, which states that every point mass attracts every other point mass by a force acting along the line intersecting both points. The force is proportional to the product of the masses and inversely proportional to the square of the distance between them:

F = G * (m₁ * m₂) / r²

where:
- F is the gravitational force between masses
- G is the gravitational constant
- m₁ and m₂ are the masses
- r is the distance between the centers of the masses

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

[MIT License](LICENSE)

## Acknowledgements

Special thanks to the Fyne toolkit developers for providing an excellent cross-platform GUI library for Go.
