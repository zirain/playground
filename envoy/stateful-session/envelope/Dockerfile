
# Start from the official Golang image for building
FROM golang:1.24 AS builder

WORKDIR /app

COPY backend.go .
# Copy go.mod and go.sum first, then download dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o backend backend.go

# Use a minimal base image for running
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/backend .

EXPOSE 8000

ENTRYPOINT ["/app/backend"]
