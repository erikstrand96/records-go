# records-go

Simple Go-based Web API.

## Features

- Written in Go
- Docker-ready build for containerized deployment

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
    - Create a `.env` file in the project root with the following variables:

      ```
      RECORDS_API_PORT=
      RECORDS_DB_NAME=
      RECORDS_DB_PASSWORD=
      RECORDS_DB_CONNECTION=
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
- `database/` - (Database logic)
- `Dockerfile`

## Author

- [erikstrand96](https://github.com/erikstrand96)
