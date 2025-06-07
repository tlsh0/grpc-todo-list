FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o user-service .

EXPOSE 50051
CMD ["./user-service"]
