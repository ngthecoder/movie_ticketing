# Movie Ticketing API
Ticketing System for Movies

## Motivation
I mainly started this project to understand how Kubernetes works in an application deployment and additionally to understand the design patterns.

## Tech Stack
Backend: Go
Database: Postgres
Infrastructure: Docker, Kubernetes
Tools: golang-migrate

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