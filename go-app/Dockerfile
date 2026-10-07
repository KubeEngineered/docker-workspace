# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency definition files first (caches layer if dependencies don't change)
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Compile static binary (disabling CGO for portability)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Stage 2: Minimal runtime stage
FROM alpine:3.19

WORKDIR /app

# Add ca-certificates in case your Go app makes HTTPS calls
RUN apk --no-cache add ca-certificates

# Copy compiled binary from builder stage
COPY --from=builder /app/main .

EXPOSE 8080

ENV PORT=8080 \
    APP_VERSION=v1.0.1

CMD ["./main"]
