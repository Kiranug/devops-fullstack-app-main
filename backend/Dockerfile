# Use official Go image as base
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go mod files for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["./main"]
