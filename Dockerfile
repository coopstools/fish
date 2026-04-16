# Stage 1: Build
FROM golang:alpine AS builder
WORKDIR /app
COPY go.* .
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/interpreter ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
# Copy the binary AND the project files (examples, etc.) from the builder
COPY --from=builder /app /app

ENTRYPOINT ["./interpreter"]