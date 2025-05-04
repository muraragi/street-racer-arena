# apps/backend/Dockerfile

# Build Stage (Production) - Keep this as is
FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .
# --- DEBUG ---
# RUN echo "--- DEBUG: Builder stage after go build ---" && ls -l /app
# RUN test -f /app/main || (echo "ERROR: /app/main NOT FOUND in builder stage!" && exit 1)
# RUN echo "--- DEBUG: /app/main confirmed FOUND in builder stage. ---"
# --- END DEBUG ---


# ---- Development Stage ----
FROM golang:1.24.2-alpine AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
COPY .air.toml .
COPY . .
EXPOSE 8080
CMD ["air"]


# ---- Final Stage (Production) ----
FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /app/main .
# --- DEBUG ---
# RUN echo "--- DEBUG: Final stage after COPY ---" && ls -l /app
# RUN test -f /app/main || (echo "ERROR: /app/main NOT FOUND in final stage after COPY!" && exit 1)
# RUN echo "--- DEBUG: /app/main confirmed FOUND in final stage. ---"
# --- END DEBUG ---
EXPOSE 8080
# Production command remains the same
CMD ["./main"]

