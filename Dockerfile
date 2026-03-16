FROM golang:1.26.1 AS builder
WORKDIR /build

COPY . /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine
WORKDIR /app

COPY --from=builder /build/main /app
EXPOSE 8080
CMD ["/app/main"]