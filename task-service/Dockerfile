FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o task-service .

EXPOSE 50052
CMD ["./task-service"]
