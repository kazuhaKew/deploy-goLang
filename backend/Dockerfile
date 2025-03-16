# Build Stage
FROM golang:1.23-alpine AS build

WORKDIR /app

# Copy the whole directory including vendor
COPY . .

# Build the application using vendor directory
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main .

# Runtime Stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from build stage
COPY --from=build /app/main .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]