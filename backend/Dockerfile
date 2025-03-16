# Build Stage
FROM golang:1.23.4-alpine AS build

WORKDIR /app

# Setup networking for go mod download
# Force IPv4 DNS resolution
RUN echo "nameserver 8.8.8.8" > /etc/resolv.conf

# Explicitly set environment variables to bypass proxy issues
ENV GOPROXY=direct
ENV GOSUMDB=off
ENV GO111MODULE=on

# Copy go.mod
COPY go.mod ./

# Try downloading dependencies with retry mechanism
RUN apk add --no-cache git ca-certificates && \
    for i in 1 2 3 4 5; do go mod download && break || sleep 5; done

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