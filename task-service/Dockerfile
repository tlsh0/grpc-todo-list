FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task-service .

FROM debian:bullseye-slim

WORKDIR /app

RUN apt-get update && apt-get upgrade -y && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/task-service .

CMD ["./task-service"]