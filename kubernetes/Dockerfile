FROM golang:1.12-alpine

# Set the Current Working Directory inside the container
WORKDIR /instagram-service

COPY go.mod ./

COPY go.sum ./

# on a production server - you can just run go mod download.
RUN go mod Download


# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . ./


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


# This container exposes port 8080 to the outside worlds
EXPOSE 8080

# Run the executable
CMD ["go","run","main.go"]

