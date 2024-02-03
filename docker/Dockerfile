# Use an official Go runtime as a parent image
FROM golang:1.18

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app
RUN go build -o main -v -gcflags "all=-N -l" -buildvcs=false ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
