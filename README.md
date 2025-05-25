# NBody-Simul: N-Body Gravitational Simulation

A gravitational simulation written in Go using the Fyne GUI toolkit.

![N-Body Simulation Demo](https://github.com/AnshJain-Shwalia/NBody-Simul/blob/main/demo.gif)

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

1. **Initialization**: The system creates celestial bodies with random positions, velocities, and masses.

2. **Force Calculation**: The program calculates the gravitational force between each pair of bodies.

3. **Position Updates**: Based on the net force, the simulation updates velocity and position vectors.

4. **Visualization**: The GUI renders each body as a circle with size proportional to its mass.

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

## Files

- `main.go`: Entry point and application setup
- `constants.go`: Simulation parameters and body structure definition
- `positionCalculation.go`: Physics calculations
- `displayHandling.go`: GUI rendering and animation
- `miscellaneous.go`: Utility functions

## License

[MIT License](LICENSE)
