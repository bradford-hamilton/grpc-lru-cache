FROM golang:alpine AS builder
WORKDIR $GOPATH/src/bradford-hamilton/grpc-lru-cache
COPY . .

# Build for linux x86
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/server .

# Expose our GRPC service
EXPOSE 21000

# Run the server binary
ENTRYPOINT ["/go/bin/server"]
