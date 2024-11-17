# Use an official Go image as the base
FROM golang:1.23-alpine

# Set environment variables (optional but recommended)
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container to handle dependencies
COPY go.mod go.sum ./

# Download all Go dependencies (this is cached)
RUN go mod download

# Copy the rest of the application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application will run on (Gin default is 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
