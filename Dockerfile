# Use an official Go runtime as a base image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /syncdeck

# Copy the entire project to the container
COPY . .

WORKDIR /syncdeck_data
RUN echo "[]" > metadata.json

WORKDIR /syncdeck/api

ENV GOPROXY=https://goproxy.io,direct

RUN go mod download

# Build the Go application
RUN go build

RUN mv api /

WORKDIR /
# Expose the port on which your Go API runs
EXPOSE 8080

# Command to run the application
CMD ["./api"]
