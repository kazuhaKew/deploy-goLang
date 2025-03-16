# Build Stage
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Runtime Stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from build stage
COPY --from=build /app/main .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]