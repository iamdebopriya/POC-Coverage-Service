# Stage 1 — Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o coverage-service .

# Stage 2 — Runtime needs BOTH Go and Node
FROM golang:1.25-alpine

# Install Node and npm
RUN apk update && apk add --no-cache nodejs npm

WORKDIR /app
COPY --from=builder /app/coverage-service .
COPY config/services.json ./config/services.json

EXPOSE 8081
CMD ["./coverage-service"]