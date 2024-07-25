# Use the official Golang image as the base image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum from the main directory to the Docker image
COPY ../../go.mod .
COPY ../../go.sum .

# Download dependencies
RUN go mod download

COPY .. .
# Copy the source code from the "cmd" folder to the Docker image
COPY cmd/main.go .

# Navigate to the directory and copy files
RUN mkdir -p internal && cp ./internal/.env .

RUN mkdir -p templates && cp ./templates/responseEmail.html .

# Compile the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a smaller base image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app

COPY --from=builder /app/internal/.env /app/internal/

COPY --from=builder /app/templates/responseEmail.html /app/templates/

# Explicitly copy the entire assets directory from the builder stage
COPY --from=builder /app/assets /app/assets

# Specify the entry point for the Docker container
ENTRYPOINT ["/app/main"]
