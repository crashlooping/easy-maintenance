# Use the official Go image as the base image
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o maintenance .

# Create a new lightweight image for serving the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/maintenance /app/maintenance

# Copy the static website content
COPY html /app/html

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the application
CMD ["/app/maintenance"]
