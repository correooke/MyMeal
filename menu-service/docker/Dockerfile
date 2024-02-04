# Stage 1: Build
FROM golang:1.21.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o category-service ./cmd/main.go 

# Stage 2: Run
FROM golang:1.21.0

WORKDIR /app

COPY --from=builder /app/category-service .

EXPOSE 8080

CMD ["./category-service"]
