# Use an official Go runtime as a base image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /syncdeck

# Copy the entire project to the container
COPY . .

RUN mkdir -p $HOME/.config/syncdeck
COPY configs/server.json /root/.config/syncdeck/server.json 

RUN mkdir $HOME/syncdeck
RUN echo "[]" > $HOME/syncdeck/metadata.json

WORKDIR /syncdeck/api
# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port on which your Go API runs
EXPOSE 5137

# Command to run the application
CMD ["./main"]
