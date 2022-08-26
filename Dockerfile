FROM golang:1.18-alpine as builder

RUN go install golang.org/dl/go1.18@latest \
  && go1.18 download

# Set the Current Working Directory inside the container
WORKDIR /build

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .
COPY src/*.go ./

RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o ./out/modmesh .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/modmesh"]