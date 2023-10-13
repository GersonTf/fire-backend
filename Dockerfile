# Stage 1: Build the Go binary
FROM golang:1.21.3-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

# Stage 2: Build a minimal image to run the binary
FROM alpine

# Copy the Go binary from the builder stage
COPY --from=builder /go/bin/app /go/bin/app

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/go/bin/app"]
