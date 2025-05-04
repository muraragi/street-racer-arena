# Build Stage
FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build .

EXPOSE 8080

RUN chmod +x ./build

CMD ["./build"]
