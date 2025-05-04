FROM golang:1.24.2-alpine AS builder
WORKDIR /app

# Only download dependencies if go.mod/go.sum change
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build statically-linked Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/server/main.go

# Minimal final image
FROM alpine:latest
WORKDIR /app

# Copy the statically-linked binary from builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./main"]
