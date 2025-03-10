# Use the official Go image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server


# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./main"]
