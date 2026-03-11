# Task API (Golang)

A minimal REST API for managing tasks built using Go.

The service allows users to:

- Create a task
- List all tasks

Tasks are stored in-memory (no database).

------------------------------------------------------------------------

# Tech Stack

-   Go 1.21+
-   Standard library (`net/http`)
-   In-memory storage
-   Unit tests
-   Graceful server shutdown

------------------------------------------------------------------------

## Project Structure

```
anra_assessment
├── cmd/server/main.go      # Application entry point
├── internal
│   ├── handler             # HTTP handlers
│   ├── service             # Business logic
│   ├── repository          # In-memory storage
│   └── model               # Task model
├── tests                   # API tests
├── go.mod
└── README.md
```

------------------------------------------------------------------------

# Prerequisites

Install Go 1.21+

Check version:
```
go version
```
------------------------------------------------------------------------

# Running the Service

Clone the repository:
```
git clone https://github.com/iabdulzahid/anra_assessment.git 
```
cd anra_assessment

Run the server:
```
go run cmd/server/main.go
```
Server starts on:
```
http://localhost:9090
```
You can also specify a custom port:
```
PORT=8080 go run cmd/server/main.go
```
------------------------------------------------------------------------

# API Endpoints

## Create Task

POST /tasks

Example request:
```
curl -X POST http://localhost:9090/tasks\
-H "Content-Type: application/json"\
-d '{"title":"Learn Go"}'
```
Example response:

201 Created
```
{ "id": "uuid", "title": "Learn Go", "status": "todo" }
```
If `status` is omitted it defaults to `todo`.

------------------------------------------------------------------------

## List Tasks

GET /tasks

Example request:
```
curl http://localhost:9090/tasks
```
Example response:

200 OK
```
[ { "id": "uuid", "title": "Learn Go", "status": "todo" }]
```
If no tasks exist the API returns:
```
[]
```
------------------------------------------------------------------------

# Running Tests

Run tests:
```
go test ./...
```
Tests use Go's `httptest` package with a table-driven testing approach.

------------------------------------------------------------------------

# Graceful Shutdown

The server supports graceful shutdown and handles:

SIGINT\
SIGTERM

Active requests are allowed to finish before shutdown.

------------------------------------------------------------------------

# Future Improvements

If more time were available:

-   Add persistent storage (PostgreSQL)
-   Add update and delete task endpoints
-   Add pagination
-   Improve test coverage
-   Add structured logging
-   Add Docker support

------------------------------------------------------------------------

# Notes

This implementation focuses on:

-   Clean project structure
-   Idiomatic Go
-   Simple REST API design
-   Clear validation and error handling
