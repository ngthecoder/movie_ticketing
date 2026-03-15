# Movie Ticketing API
Ticketing System for Movies

## Motivation
I mainly started this project to understand how Kubernetes works in an application deployment and additionally to understand the design patterns.

## Tech Stack
Backend: Go
Database: Postgres
Infrastructure: Docker, Kubernetes
Tools: golang-migrate

## Structure
The bookings/movies/payments/screenings/theaters/users folders each has
- handler.go: Contains HTTP handlers
- service.go: Contains internal processes
- repository.go: Contains db operations

The db folder contains the db management functions and migrations folder has necessary SQL operations.

The root folder has .env (defines environment variables), docker-compose.yml (sets up Postgres container), go.mod/go.sum (manage dependencies), main.go (orchestrator for the HTTP server)

## Endpoints
GET /movies
GET /movies/:id

## Getting Started
### Prerequisites
- Docker
- Go 1.21+

### Run Locally
```bash
docker compose up -d
migrate -path migrations -database "postgres://..." up
go run main.go
```