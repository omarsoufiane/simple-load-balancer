# Use the official Go image to build the application
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code and necessary files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
