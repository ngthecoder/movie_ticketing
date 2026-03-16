# Movie Ticketing API
Ticketing System for Movies

## Motivation
I mainly started this project to understand how Kubernetes works in an application deployment and additionally to understand the design patterns.

## New Lessons
### PostgreSQL Transaction
I mostly used independent queries in movies, theaters, etc but in the bookings package, I used `db.Beginx()` to achieve Atomicity (one of the ACID) when executing two separate queries (creating a booking and updating the number of available seats).

## Intentional Designs
### Repository Pattern
The app follows a handler -> service -> repository -> DB layered structure. Each layer only communicates with the layer directly below it, hiding database implementation details from the upper layers. Swapping out the database only requires changes in the repository layer.

### Sentinel Errors
Services translate low-level SQL errors (e.g. `sql.ErrNoRows`) into domain-specific errors (e.g. `ErrNotFound`). This ensures handlers only deal with application-level errors and remain unaware of database internals.

### Password Hashing
Passwords are hashed using bcrypt before being stored in the database. Even if the database is breached, the original passwords cannot be recovered.

### Dependency Injection
All dependencies are assembled in `main.go` in order: repository → service → handler. Each layer declares what it needs through its constructor, keeping the dependency relationships explicit and centralized.

## Tech Stack
- Backend: Go
- Database: Postgres
- Infrastructure: Docker, Kubernetes
- Tools: golang-migrate

## Structure
The bookings/movies/payments/screenings/theaters/users folders each has
- handler.go: Contains HTTP handlers
- service.go: Contains internal processes
- repository.go: Contains db operations

The db folder contains the db management functions and migrations folder has necessary SQL operations.

The root folder has .env (defines environment variables), docker-compose.yml (sets up Postgres container), go.mod/go.sum (manage dependencies), main.go (orchestrator for the HTTP server)

## Endpoints
- GET /movies
- GET /movies/:id
- GET /theaters
- GET /theaters/:id
- POST /users/register
- POST /users/login
- GET /screenings
- GET /screenings/:id
- POST /bookings
- POST /payments

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