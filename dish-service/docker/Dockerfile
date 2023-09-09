# Stage 1: Build
FROM golang:1.21.0 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Fetch dependencies first; they are less susceptible to change on every build 
# and will therefore be cached for speeding up the next build
RUN go mod download

# Copy the local package files to the container’s workspace.
COPY . ./

# Build the application.
RUN go build -o dish-service ./cmd/main.go 

# Stage 2: Run
FROM golang:1.21.0

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/dish-service .

# Expose port 8080
EXPOSE 8080

# Command to run the application when the docker container starts
CMD ["./dish-service"]
