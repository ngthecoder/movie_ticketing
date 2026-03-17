# Movie Ticketing API
A RESTful API for movie ticketing built with Go, PostgreSQL, and Kubernetes.
Designed to explore real-world application deployment on Kubernetes and backend design patterns including layered architecture, sentinel errors, and database transactions.

## New Lessons
### PostgreSQL Transaction
I mostly used independent queries in movies, theaters, etc but in the bookings package, I used `db.Beginx()` to achieve Atomicity (one of the ACID) when executing two separate queries (creating a booking and updating the number of available seats).

### Environment Variables Injection in Kubernetes
There are two ways to inject environment variables into a Deployment: `envFrom` and `env`. I used `envFrom` in `app-deployment.yaml` because the key names in ConfigMap match exactly what the app expects. I used `env` in `postgres-deployment.yaml` because the Postgres image expects different variable names (e.g. `POSTGRES_USER`) than what is defined in ConfigMap (e.g. `DB_USER`), so each variable needs to be mapped individually.

### Use of Labels and Metadata
K8s service discovery works in two steps.

First, the app pod looks for `postgres-service` by name (`DB_HOST` in `configmap.yaml`). K8s resolves this name as DNS within the cluster, so `DB_HOST` must match `metadata.name` in `postgres-service.yaml`.

Second, once the postgres Service receives a request, it forwards it to the matching pod using `selector.app`, which must match `spec.template.metadata.labels.app` in `postgres-deployment.yaml`.

### Use of LivenessProbe and ReadinessProbe
In `app-deployment.yaml`, I used both LivenessProbe and ReadinessProbe and I chose /ping for LivenessProbe and /movies for ReadinessProbe because successful check against /ping only proves the app pod is alive and doesn't indicate whether the pod is properly talking with the postgres pod. On the other hand, successful check against /movies shows that the app pod can retrieve data from the postgres pod and it is ready to serve the application.

### AccessMode for Persistent Volume Claim
In `postgres-pvc.yaml`, I set `ReadWriteOnce` for `accessModes`, which means the volume can only be mounted by a single node at a time. This is appropriate because there is only one postgres pod in this deployment. If multiple pods needed to share the same volume simultaneously, `ReadWriteMany` would be required instead.

### Use of Init Container
After applying the manifests, the app pods were failing the readiness check because the postgres pod had no tables yet. The workaround was to manually port-forward to the postgres pod and run the migration from my local machine. To fix this, I created `Dockerfile.migrate` which copies the migrations folder into the `migrate/migrate` image, and added an init container to the app deployment that runs the migration before the app container starts.

### SELECT FOR UPDATE and CHECK Constraint
When creating a booking, the available seats check and the UPDATE ran in separate steps, leaving a window where concurrent transactions could both pass the check before either updated the seats, a race condition. To fix this, I moved the check inside the transaction in `bookings/repository.go` and used `SELECT FOR UPDATE` to lock the row before checking and updating, ensuring only one transaction can proceed at a time.

I also added `CHECK (available_seats >= 0)` to the screenings table as a last line of defense at the DB level.

## Intentional Designs
### Repository Pattern
The app follows a handler -> service -> repository -> DB layered structure. Each layer only communicates with the layer directly below it, hiding database implementation details from the upper layers. Swapping out the database only requires changes in the repository layer.

### Sentinel Errors
Services translate low-level SQL errors (e.g. `sql.ErrNoRows`) into domain-specific errors (e.g. `ErrNotFound`). This ensures handlers only deal with application-level errors and remain unaware of database internals.

### Password Hashing
Passwords are hashed using bcrypt before being stored in the database. Even if the database is breached, the original passwords cannot be recovered.

### Dependency Injection
All dependencies are assembled in `main.go` in order: repository → service → handler. Each layer declares what it needs through its constructor, keeping the dependency relationships explicit and centralized.

### Kubernetes Service Types
I used NodePort for service type in `app-service.yaml` and ClusterIP in `postgres-service.yaml` because the postgres pod should only be accessible within the cluster while the app/API pod needs to be accessed from the public.

### Kubernetes Secret Encoding
Credentials stored in Kubernetes Secrets must be base64 encoded. Note that base64 is not encryption and it can be easily decoded. Secrets are simply base64 encoded to ensure safe handling of binary data in YAML format because credentials often contain special characters like # and ! that could break YAML parsing. For production use, additional security measures such as encrypted Secret stores are recommended.

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
- Go 1.25+

#### Commands
```bash
docker compose up --build
```

### Run with Minikube
#### Prerequisites
- Docker
- Go 1.25+
- Minikube

#### Commands
**Terminal 1**
```bash
minikube start
eval $(minikube docker-env)
docker build -t movie_ticketing_app:v1 .
docker build -f Dockerfile.migrate -t movie_ticketing_migrate:v1 .
kubectl apply -f k8s/

# Keep this running and move to Terminal 2
kubectl port-forward pod/movie-ticketing-postgres-deployment-<id> 5432:5432
```
**Terminal 2**
```bash
# Inject seed data
psql "postgres://postgres:<password>@localhost:5432/movie_ticketing?sslmode=disable" -f db/seed.sql  
```
**Back to Terminal 1 (after stopping port-forward with Ctrl+C)**
```bash
minikube service app-service --url   
```