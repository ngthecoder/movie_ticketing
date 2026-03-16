# Movie Ticketing API
A RESTful API for movie ticketing built with Go, PostgreSQL, and Kubernetes.
Designed to explore real-world application deployment on Kubernetes and backend design patterns including layered architecture, sentinel errors, and database transactions.

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
| Category | Technology |
|----------|------------|
| Backend | Go, Gin |
| Database | PostgreSQL |
| Infrastructure | Docker, Kubernetes (minikube) |
| Tools | golang-migrate, sqlx |

## Structure
```
movie_ticketing/
├── bookings/
│   ├── handler.go       # HTTP handlers
│   ├── service.go       # Business logic
│   └── repository.go    # DB operations
├── movies/              # Same structure as bookings
├── payments/            # Same structure as bookings
├── screenings/          # Same structure as bookings
├── theaters/            # Same structure as bookings
├── users/               # Same structure as bookings
├── db/
│   ├── db.go            # DB connection
│   └── seed.sql         # Seed data
├── migrations/          # SQL migration files
├── k8s/                 # Kubernetes manifests
├── docker-compose.yml   # Local Postgres setup
├── Dockerfile           # Multi-stage build
├── main.go              # Entry point, dependency wiring
└── go.mod
```

## Endpoints
| Method | Path | Description |
|--------|------|-------------|
| GET | `/movies` | Returns all movies |
| GET | `/movies/:id` | Returns one movie by ID |
| GET | `/theaters` | Returns all theaters |
| GET | `/theaters/:id` | Returns one theater by ID |
| POST | `/users/register` | Registers a user, returns user info |
| POST | `/users/login` | Logs in a user, returns user info |
| GET | `/screenings` | Returns all screenings |
| GET | `/screenings/:id` | Returns one screening by ID |
| POST | `/bookings` | Creates a booking, returns booking info |
| POST | `/payments` | Confirms a booking, returns payment info |

## Getting Started
### Run with Docker Compose
#### Prerequisites
- Docker
- Go 1.21+

#### Commands
```bash
docker compose up --build
```

### Run with Minikube
#### Prerequisites
- Docker
- Go 1.21+
- Minikube

#### Commands
**Termianl 1**
```bash
minikube start
eval $(minikube docker-env)
docker build -t <image name>:<tag> .
kubectl apply -f k8s/

# Keep this running and move to Terminal 2
kubectl port-forward pod/movie-ticketing-postgres-deployment-<id> 5432:5432
```
**Terminal 2**
```bash
migrate -path migrations -database "postgres://postgres:<password>@localhost:5432/movie_ticketing?sslmode=disable" up
psql "postgres://postgres:<password>@localhost:5432/movie_ticketing?sslmode=disable" -f db/seed.sql  
```
**Back to Termianl 1 (after stopping port-forward with Ctrl+C)**
```bash
minikube service app-service --url   
```