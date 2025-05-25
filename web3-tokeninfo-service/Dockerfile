# Use the official Golang image as the base image
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application for Linux architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o token-info cmd/app/main.go

# Use a minimal base image for the final container
FROM alpine:3.21

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/token-info .

# Expose the application port
EXPOSE 8082

# Command to run the application
CMD ["./token-info"]
