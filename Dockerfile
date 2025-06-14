# 1. Use the official Go image as a builder stage 1.1. AS builder: The stage's name (for multi-stage build)
FROM golang:1.24.4-alpine3.22 AS builder

# 2. Set the working directory inside the container
WORKDIR /app

# 3. Copy the application source code
COPY app/ .

# 4. Download dependencies
RUN go mod download

# 5. Build the Go binary
RUN go build -o myfoodapplication_homelab

# 6. Final stage: use minimal base image to run the binary
FROM alpine:latest

# 7. Create working directory
WORKDIR /app

# 8. Copy the binary from builder stage
COPY --from=builder /app/myfoodapplication_homelab .

# 9. Document the port the app listens on
EXPOSE 8080

# 10. Set default command to run
CMD ["./myfoodapplication_homelab"]
