# --- Build stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for Go modules if needed
RUN apk add --no-cache git

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o users ./cmd/users

# --- Final stage ---
FROM alpine:3.21

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/users .

# Copy config file if needed (optional)
# COPY config.yml .

# Set environment variables if needed
# ENV NATS_URL=nats://nats:4222

# Run the binary
ENTRYPOINT ["./users"]
