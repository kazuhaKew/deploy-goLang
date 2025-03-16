FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /app

# Add CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create logs directory with proper permissions
RUN mkdir -p /app/logs && chmod 755 /app/logs

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["./main"]