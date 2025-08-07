# Use official Go 1.24 image
FROM golang:1.24-alpine

# Install Git (in case some deps require it)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the app from cmd/
RUN go build -o main ./cmd

# Expose port
EXPOSE 5000

# Run the binary
CMD ["./main"]
