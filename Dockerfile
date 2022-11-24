FROM golang:1.12-alpine

# Set the Current Working Directory inside the container
WORKDIR /instagram-service

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . /instagram-service/

# on a production server - you can just run go mod download.
RUN go mod Download


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go","run","main.go"]