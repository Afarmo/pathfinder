# 🚆 Stations Pathfinder
![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
## Overview

Stations Pathfinder is a Go-based train dispatching simulation that routes multiple trains through a railway network while minimizing the total number of turns required to reach their destination.

The program parses a railway network map, validates its structure, discovers multiple viable routes between two stations, distributes trains across those routes, and outputs turn-by-turn movement instructions.

---

## Setup

### Install & Run

Clone the repository:

```bash
git clone https://github.com/Afarmo/pathfinder.git
cd pathfinder
```

Run the program:

```bash
go run ./cmd [path to file containing network map] [start station] [end station] [number of trains]
```

Example:

```bash
go run ./cmd maps/network.map waterloo st_pancras 4
```

---

## Usage

Given the following railway network:

```text
stations:
waterloo,3,1
victoria,6,7
euston,11,23
st_pancras,5,15

connections:
waterloo-victoria
waterloo-euston
st_pancras-euston
victoria-st_pancras
```

Running:

```bash
go run ./cmd maps/network.map waterloo st_pancras 4
```

Produces:

```text
T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras
```

Each line represents a turn.

Each train movement is displayed in the format:

```text
T<train-id>-<station>
```

---

## Features

### Network Map Parsing

Parses railway network definitions consisting of:

- Stations
- Coordinates
- Connections

### Validation

Validates command-line arguments and network maps before pathfinding begins.

Supported validation includes:

- Missing or excess arguments
- Invalid file paths
- Invalid train counts
- Non-existent start stations
- Non-existent end stations
- Identical start and end stations
- Duplicate station names
- Duplicate station coordinates
- Invalid coordinate values
- Invalid station names
- Duplicate connections
- Connections referencing unknown stations
- Missing `stations:` section
- Missing `connections:` section
- Malformed map syntax
- Networks containing more than 10,000 stations
- No available path between the selected stations



## Pathfinding Strategy

### Multiple Path Discovery

Discovers multiple valid routes between the start and destination stations while respecting station occupancy constraints.

### Breadth-First Search (BFS)

Breadth-First Search is used to find augmenting paths through the residual network.

Because BFS explores nodes level-by-level, it guarantees the shortest augmenting path in terms of edge count during each Edmonds-Karp iteration.

### Edmonds-Karp Maximum Flow

The railway network is transformed into a flow network where intermediate stations are split into inbound and outbound nodes.

This allows the algorithm to enforce station-capacity constraints while searching for multiple non-overlapping routes.

The Edmonds-Karp algorithm repeatedly:

1. Uses BFS to find an augmenting path.
2. Pushes flow through the discovered path.
3. Updates the residual network.
4. Repeats until no additional augmenting paths exist.

The resulting flow network is then used to extract all viable train routes.

---

### Train Scheduling

Distributes trains across the available routes and generates optimized turn-by-turn movement instructions.

The scheduler:

- Assigns trains to discovered paths
- Balances path utilization
- Reduces congestion
- Minimizes overall completion time
- Generates movement instructions for every turn

---

## Output

The program outputs train movements turn-by-turn.

Example:

```text
T1-victoria T2-euston
T1-st_pancras T2-st_pancras T3-victoria T4-euston
T3-st_pancras T4-st_pancras
```

Where:

- `T1`, `T2`, `T3`, etc. represent train identifiers.
- The station following the dash represents the train's position after that turn.

---

## Project Structure

```text
.
├── cmd
│   └── main.go
├── internal
│   ├── cli
│   │   └── args.go
│   ├── models
│   │   └── structs.go
│   ├── parser
│   │   ├── parser.go
│   │   └── validation.go
│   ├── pathfinder
│   │   ├── pathFinder.go
│   │   └── bfs.go
│   │   └── maxFlow.go
│   │   └── graph.go
│   │   └── extract.go
│   └── scheduler
│       └── scheduler.go
├── maps
│   └── *.map
├── go.mod
└── README.md
```