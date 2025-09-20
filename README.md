# pprotein

[![Go Version](https://img.shields.io/badge/go-1.23.3-blue.svg)](https://golang.org/)

**pprotein** is a comprehensive performance analysis tool designed specifically for ISUCON competitions and high-performance web applications. It provides real-time profiling, log analysis, and performance monitoring with an intuitive web interface.

## 🚀 Features

### Core Functionality
- **Real-time Performance Profiling** - CPU, memory, and goroutine analysis using Go pprof
- **HTTP Log Analysis** - Integrated with `alp` for access log analysis
- **MySQL Slow Query Analysis** - Built-in slow query log processor
- **Git Repository Integration** - Track performance changes with commit information
- **Real-time Updates** - WebSocket-based live data streaming
- **Group Collaboration** - Share profiling results across team members

### Web Framework Integration
- **Echo v3/v4** - Seamless integration with Echo framework
- **Gin** - Native Gin middleware support
- **Gorilla Mux** - Direct router integration
- **Standalone** - Works with any Go web application

### User Interface
- **Modern Web UI** - Vue.js 3 + TypeScript frontend
- **Interactive Charts** - Real-time performance visualization
- **Responsive Design** - Works on desktop and mobile
- **Dark/Light Mode** - Comfortable viewing in any environment

## 📦 Installation

### Prerequisites
- Go 1.23.3 or later
- Node.js 18+ and npm (for frontend development)

### Quick Start

```bash
# Clone the repository
git clone https://github.com/mism-mism/pprotein.git
cd pprotein

# Build the application
make build

# Run the main server
./pprotein
# or
make run

# Run the agent (for remote profiling)
./pprotein-agent --git-dir /path/to/your/project
# or
make run-agent
```

### Using Docker

```bash
# Run with Docker Compose (includes mock environment)
docker-compose up

# Access the web interface
open http://localhost:9000
```

## 🛠️ Usage

### Basic Setup

#### 1. Main Server
The main pprotein server provides the web interface and data collection:

```bash
# Default setup (port 9000)
./pprotein

# Custom port
PORT=8080 ./pprotein
```

#### 2. Agent Mode
Use pprotein-agent for remote profiling of your applications:

```bash
# Basic usage
./pprotein-agent

# Specify git repository and port
./pprotein-agent --git-dir /path/to/your/project --port 19000

# Show help
./pprotein-agent --help
```

### Framework Integration

#### Echo v4 Integration
```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/kaz/pprotein/integration/echov4"
)

func main() {
    e := echo.New()
    echov4.Integrate(e)
    
    // Your routes here
    e.GET("/", handler)
    
    e.Start(":8080")
}
```

#### Gin Integration
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/kaz/pprotein/integration/gin"
)

func main() {
    r := gin.Default()
    gin.Integrate(r)
    
    // Your routes here
    r.GET("/", handler)
    
    r.Run(":8080")
}
```

#### Standalone Integration
```go
package main

import (
    "net/http"
    "github.com/kaz/pprotein/integration/standalone"
)

func main() {
    // Start pprotein agent in background
    go standalone.Integrate(":19000")
    
    // Your application
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Main server port | `9000` |
| `PPROTEIN_HTTPLOG` | HTTP access log path | `/var/log/nginx/access.log` |
| `PPROTEIN_SLOWLOG` | MySQL slow query log path | `/var/log/mysql/mysql-slow.log` |
| `PPROTEIN_GIT_REPOSITORY` | Git repository path | `.` (current directory) |

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Web Frontend  │    │  Main Server    │    │     Agent       │
│   (Vue.js)      │◄──►│   (Go Echo)     │◄──►│  (Standalone)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌─────────────────┐    ┌─────────────────┐
                       │    Storage      │    │   Target App    │
                       │   (BoltDB)      │    │  (Your Service) │
                       └─────────────────┘    └─────────────────┘
```

### Core Components

- **`cli/pprotein`** - Main web server with UI
- **`cli/pprotein-agent`** - Lightweight profiling agent
- **`cli/pprotein-mock`** - Mock server for testing
- **`internal/collect`** - Data collection and processing
- **`internal/storage`** - Persistent data storage (BoltDB)
- **`internal/event`** - Real-time WebSocket communication
- **`view/`** - Vue.js frontend application

## 🔍 ISUCON Usage Patterns

### 1. Baseline Measurement
```bash
# Start profiling before optimization
./pprotein-agent --git-dir /home/isucon/webapp

# Run benchmark and collect data
# Access http://localhost:9000 to view results
```

### 2. Continuous Monitoring
```bash
# Profile during development
make run &
./pprotein-agent --port 19001 &

# Monitor changes in real-time
```

### 3. Team Collaboration
- Use group functionality to share results
- Track performance across git commits
- Compare optimization attempts

## 🛡️ Performance Analysis Features

### CPU Profiling
- Hotspot identification
- Function-level analysis
- Goroutine monitoring
- Memory allocation tracking

### HTTP Log Analysis
- Endpoint performance ranking
- Response time distribution
- Request volume analysis
- Error rate monitoring

### MySQL Analysis
- Slow query identification
- Execution frequency analysis
- Index usage optimization
- Query performance trends

## 🔧 Development

### Cursor Integration

This project includes custom Cursor commands for streamlined development workflow. Use `/` in Cursor Agent to access these commands:

- **`/build`** - Build all binaries with frontend
- **`/test`** - Run comprehensive test suite
- **`/lint`** - Execute linters and code quality checks
- **`/start-dev`** - Start development environment
- **`/deploy`** - Deploy to production environment
- **`/profile-isucon`** - Start ISUCON performance analysis
- **`/create-pr`** - Create high-quality pull requests

Commands are stored in `.cursor/commands/` and follow the patterns established in the [Cursor changelog](https://cursor.com/ja/changelog).

### VS Code/Cursor Tasks

Use Ctrl+Shift+P (Cmd+Shift+P on Mac) and search for "Tasks: Run Task" to access:

- **Build All** - Complete project build
- **Run Main Server** - Start pprotein server
- **Run Agent** - Start profiling agent  
- **Build Frontend** - Vue.js frontend build
- **Test Go** - Go unit tests with race detection
- **Lint Go/Frontend** - Code quality checks

### Building from Source

```bash
# Install dependencies
go mod download
npm --prefix view install

# Build frontend
npm --prefix view run build

# Build binaries
make build

# Clean build artifacts
make clean
```

### Running in Development Mode

```bash
# Backend
make run

# Frontend (in separate terminal)
cd view
npm run dev
```

### Testing

```bash
# Run Go tests
go test ./...

# Run frontend tests
cd view
npm test
```

## 📊 Monitoring Endpoints

When integrated, pprotein exposes these debug endpoints:

- `/debug/pprof/` - Standard Go pprof endpoints
- `/debug/fgprof` - fgprof CPU profiling
- `/debug/log/httplog` - HTTP access log tail
- `/debug/log/slowlog` - MySQL slow query log tail

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow the established project structure
- Add tests for new functionality
- Update documentation as needed
- Use the provided Cursor Rules for code quality

## 📄 License

This project is a fork and does not include a license. Please check with the original repository for licensing information.

## 🙏 Acknowledgments

- Built for the ISUCON community
- Inspired by the need for better performance analysis tools
- Thanks to all contributors and users

## 📞 Support

- GitHub Issues: [Report bugs or request features](https://github.com/mism-mism/pprotein/issues)
- ISUCON Community: Share your experiences and get help

---

**Happy Profiling! 🚀**
