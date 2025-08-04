# Go-Go-Go

A full-stack multiplayer [Alkkagi (Stone-Shooting)](https://en.wikipedia.org/wiki/Alkkagi) game built entirely in Go, featuring a `terminal user interface (TUI)`, `real-time WebSocket communication` and `game logic with physics`.

<img src="./docs/images/mainPage.png" alt="Game Screenshot" width="600">
<img src="./docs/images/play-preview.gif" alt="Game Screenshot" width="600">

## Table of Contents
- [Project Overview](#project-overview)
- [How to Play](#how-to-play)
- [Installation](#installation)
- [Core Components](#core-components)

<br>

## Project Overview
### 🖥️ **Terminal User Interface (TUI)**
- **Modern TUI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework for a gui-like experience in the terminal.
- **Real-time Rendering**: Render game state and animations by using handling custom animation logic.

### 🌐 **Real-time Multiplayer**
- **WebSocket Communication**: Instant game state synchronization by sending game events and animation by using [Gorilla WebSocket](https://github.com/gorilla/websocket).
- **Live Message System**: In-game messaging for player interaction. Messages include chat and game events.

### 🎮 **Game Mechanics**
- **Physics Engine**: Realistic stone collision and movement simulation with calculations for velocity and acceleration.
- **Relative Positioning**: Stones are positioned in game board (100x100) with relative coordinates and rendered in user's current terminal window size.

### 💾 **Recording History**
- **Game Records**: Automatic recording of all game sessions with saving the event history. ([postgreSQL](https://www.postgresql.org))

<br>

## How to Play
1. **Start**: Select the start menu to match with another player. Use arrow keys or `H`, `J`, `K`, `L` keys to navigate. (see below for help)
<br><img src="./docs/images/selectMenu.gif" alt="Game Screenshot" width="600">

2. **Match**: Wait for another player to join your session. when both players are ready, the game starts automatically.
<br><img src="./docs/images/multi.gif" alt="Game Screenshot" width="600">

3. **Game**: Use arrow keys or `H`, `J`, `K`, `L` keys to navigate and **control** your shoot. Press `I` to **chat** with other player. You can **resize** the window to see the game board and stones in more detail.

4. **Win**: If one player has no more stones left, the game ends and exit with an announcement of the winner.

<br>

### Previews

#### ***Shoot***
---
control your shoot's direction and power with arrow keys or `H`, `J`, `K`, `L` keys.

<img src="./docs/images/manipulation.gif" alt="Game Screenshot" width="600">

---

when you shoot, the stone will move with a physics simulation, including friction and collision detection. Make other player's stones out of the board!

<img src="./docs/images/preview_collision.gif" alt="Game Screenshot" width="600">

#### ***Chat*** 
---
You can chat with other player or receive game events message from the server.

<img src="./docs/images/chatting.gif" alt="Game Screenshot" width="600" >

#### ***Resize***
---
To see the game board and stones in more real shape, you can resize the window.
<img src="./docs/images/resizing.gif" alt="Game Screenshot" width="600">

<br>

## Installation

#### Prerequisites
- [Go](https://go.dev/doc/install) (version 1.24.4 or later)

#### 1. Clone the repository
```bash
git clone https://github.com/yanmoyy/go-go-go.git
```

#### 2. Install dependencies
```bash
cd go-go-go
go mod download
```

*By default, the server and client use `.env` file for configuration. You can use the example file already in the repository.*
```bash
# Server
SERVER_PORT="8080"                # server port

# Client
HTTP_BASE="http://localhost:8080" # base URL for HTTP requests
WS_BASE="ws://localhost:8080"     # base URL for WebSocket connections
```

#### 3. Run the server
```bash
go run ./cmd/server
```
If you want to save the game history to database, you need to specify the database configuration in `.env` file. And run with `-db` flag.

```bash
# Database (optional)
DB_HOST="localhost"               # database host
DB_PORT="5432"                    # database port
DB_USER="postgres"                # database user
DB_PASSWORD="strongpassword"      # database password
DB_NAME="go-go-go"                # database name
```
```bash
go run ./cmd/server -db
```

#### 4. Run the client
```bash
go run ./cmd/client
```

<br>

## Core Components

>We tried to build this project from scratch, using as few dependencies as possible to learn Go deeply. Below are the main components of the project:

- **Logging**: Using [Slog](https://pkg.go.dev/log/slog), a structured logger for Go, to debug the application, track game flow, and resolve issues.
- **Testing**: Go’s testing environment enabled us to write testable code and address complex challenges (e.g., physics and UI rendering). We used [Require](https://pkg.go.dev/github.com/stretchr/testify/require) for assertions.
- **Websocket**: Using [Gorilla WebSocket](https://github.com/gorilla/websocket) for real-time communication. We leveraged goroutines for concurrent connections and channels for message passing, following the library’s documentation and examples.
- **TUI**: Using [Bubble Tea](https://github.com/charmbracelet/bubbletea) for a modern TUI framework, styled with [Lipgloss](https://github.com/charmbracelet/lipgloss) and enhanced with [Bubbles](https://github.com/charmbracelet/bubbles) for UI components. We implemented custom UI layout components (e.g., Column, Row, FlexContainer) similar with Flutter widgets. Also, the TUI structure was inspired by [opencode](https://github.com/opencode-ai/opencode) project.
- **Responsive UI**: The game board is rendered using relative coordinates, adapting to the user’s current terminal window size. Run the application and resize your terminal to see the UI adjust dynamically!
- **Physics Simulation**: A custom 2D physics engine for realistic stone movement and collision detection, built entirely in Go without external libraries. (That’s why we named this project “Go-Go-Go”!)

### Future Plans
- Single-player mode with AI opponent
- User customization options (e.g., stone colors, board size)
- Special stone types (e.g., stones with special abilities)
- User Accounts and Authentication

### Contributors
- [@yanmoyy](https://github.com/yanmoyy)
- [@Gutssssssssss](https://github.com/Gutsssssssssss)

### License
[MIT](./LICENSE)
