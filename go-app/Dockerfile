# --- Stage 1: Build Stage ---
FROM golang:1.22-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go module files first to leverage Docker layer caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build a static Go binary targeting Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# --- Stage 2: Minimal Runtime Stage ---
FROM alpine:3.19

# Install ca-certificates in case your Go app makes outbound HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# Set default environment variables (can be overridden at runtime)
ENV PORT=8080 \
    APP_VERSION=v3.0.0 \
    APP_ENV=production

# Expose the application port
EXPOSE 8080

# Configure a Docker container health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./main"]
