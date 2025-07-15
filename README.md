# records-go

`records-go` is a simple Go-based HTTP API server template, designed for fast and easy deployment of basic APIs. It demonstrates a modular structure, Docker integration, and basic environment variable configuration.

## Features

- Written in Go
- Modular code structure (API, database, main entry)
- Loads environment variables using `godotenv`
- Configurable port via `RECORDS_API_PORT` environment variable
- Basic HTTP route (`/`) returning "Hello World!"
- Docker-ready build for containerized deployment

## Usage

The main application spins up an HTTP server, routing requests to the base path (`/`) which responds with a simple message.

## Installation

### Prerequisites

- Go 1.24+
- Docker (optional, for containerization)

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/erikstrand96/records-go.git
   cd records-go
   ```

2. **Configure Environment**
   - Create a `.env` file in the project root.
   - Add a port variable:
     ```
     RECORDS_API_PORT=
     ```

3. **Run Locally**
   ```bash
   go run main.go
   ```
   The server will start on the port specified in `.env`.

4. **Run with Docker**
   ```bash
   docker build -t records-go .
   docker run -p hostport:containerport records-go
   ```

## Folder Structure

- `main.go` - Application entrypoint
- `api/` - API routing and server logic
- `database/` - (Placeholder for database logic)
- `Dockerfile` - Containerization script

## Author

- [erikstrand96](https://github.com/erikstrand96)

---
For questions or contributions, open an issue or PR!
