FROM golang:1.19.3-alpine

# Set the Current Working Directory inside the container
WORKDIR /instagram-service

COPY go.mod ./

COPY go.sum ./

# Go modules wiil be installed into a directory inside the image
RUN go mod Download


# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . ./


# The result of that command will be a static application binary named instagram-app and located in the root of the filesystem of the image that we are building.
RUN go build -o /instagram-app


# This container exposes port 8080 to the outside worlds
EXPOSE 8080

# Run the executable
CMD ["go","run","main.go"]

