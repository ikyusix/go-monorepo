# Dockerfile for user-service

# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download and cache Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o user-service

# Expose port 8081 to the outside world
EXPOSE 8081

# Run the Go app
CMD ["./user-service"]
