# Building SteamInputDB

## Prerequisites

### System Requirements

- Go 1.26 or later
- Node.js 25+
- PostgreSQL 18+
- Protocol Buffers Compiler (protoc)
- _Optional(!):_ Docker and Docker Compose (for containerized deployment)

### Package Installation

**Arch Linux:**

```bash
sudo pacman -S go nvm postgresql docker docker-compose protobuf
```

**Ubuntu:**

```bash
sudo apt install golang nvm postgresql docker.io docker-compose protobuf-compiler
```

### Go Tools

Install required Go tools:

```bash
# Protocol Buffers Go plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Air for hot reloading
go install github.com/air-verse/air@latest

# golangci-lint for linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## Building the Project

### Initialize Submodules

The project uses Steam protocol buffer definitions as a submodule:

```bash
git submodule update --init --recursive
```

### Generate Protocol Buffer Types

Generate Go and TypeScript types from Steam protobufs:

**Backend:**

```bash
cd backend
make gen
```

**Frontend:**

```bash
cd frontend
npm run gen-proto
```

### Backend

**Install dependencies:**

```bash
cd backend
make deps
```

**Build the binary:**

```bash
make build
```

The binary will be located at `backend/dist/steaminputdb`.

**Run tests:**

```bash
make test
```

**Development with hot reload:**

```bash
make dev
```

### Frontend

**Development server:**

```bash
cd frontend
npm run dev
```

**Build for production:**

```bash
npm run build
```

**Run tests:**

```bash
npm test
```

## Running with Docker Compose

### Development Environment

```bash
docker compose -f compose.dev.yml up
```

### Production Environment

```bash
docker compose up -d
```

The services will be available at:

- Backend: `http://localhost:8888`
- Backend API: `http://localhost:8889`
- Metrics: `http://localhost:8899`

## Database Setup

The backend automatically runs migrations on startup.

## Configuration

Backend configuration is managed via TOML files. See `docker/steaminputdb/config.toml` for an example.

Environment variables can also be used:

- `LISTEN_ADDRESS` - Main server listen address (default: `:8888`)
- `API_LISTEN_ADDRESS` - API server listen address (default: `:8889`)
- `METRICS_LISTEN_ADDRESS` - Metrics server listen address (default: `:8899`)
- `DATABASE_URL` - PostgreSQL connection string

## Makefile Targets

Run `make help` in the backend directory to see all available targets:

```bash
cd backend
make help
```

Common targets:

- `make build` - Build the binary
- `make test` - Run tests
- `make dev` - Run with hot reload
- `make lint` - Run linter
- `make fmt` - Format code
- `make gen` - Generate protobuf types
- `make clean` - Remove build artifacts
