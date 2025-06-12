FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go build -o apigateway-service .

EXPOSE 8080
CMD ["./apigateway-service"]
