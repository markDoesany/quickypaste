FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy Go modules files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the entire backend code
COPY . .

# Set working directory for the build process
WORKDIR /app/backend/cmd/quickypaste

# Build the Go application
RUN go build -o /app/main .

# Use a minimal runtime image
FROM alpine:latest

# Set working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
COPY backend/cmd/quickypaste/.env .env

# Expose the application port (change if needed)
EXPOSE 8080

# Run the application
CMD ["./main"]
