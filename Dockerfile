# Use the official Golang image as a base
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -a -v -o /app/dist/main .

# Use a minimal Alpine image for the final stage
FROM alpine:3.17

# Copy the binary from the builder stage
COPY --from=builder /app/dist/main /usr/local/bin/code-roasting

# Set the entry point
ENTRYPOINT ["code-roasting"]