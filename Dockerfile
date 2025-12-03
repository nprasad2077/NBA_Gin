# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary. -ldflags="-w -s" strips debug info for smaller size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api

# Run Stage
FROM alpine:latest

WORKDIR /app

# Install certificates for HTTPS calls if needed
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/.env . 
# Note: In production, .env usually injected by orchestrator, but copying for simplicity here

EXPOSE 8080

CMD ["./main"]