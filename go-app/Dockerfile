# ================================
# Stage 1: Build Stage
# ================================
FROM golang:1.22-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker layer caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application as a static binary
# CGO_ENABLED=0 disables C dependencies to ensure portability
# GOOS=linux targets Linux OS
# -ldflags="-w -s" strips debug information to minimize binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# ================================
# Stage 2: Minimal Runtime Stage
# ================================
FROM alpine:3.19 AS runner

# Create a non-root user and group for security best practices
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# Change ownership of the binary to non-root user
RUN chown appuser:appgroup /app/main

# Switch to non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]