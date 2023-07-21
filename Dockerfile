# Start from a base image with Go installed
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o dekube

# Expose a port if your Go application listens on a specific port
EXPOSE 8080

# Set the entry point for the container
CMD ["./dekube"]
