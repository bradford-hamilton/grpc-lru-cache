FROM golang:alpine AS builder
WORKDIR $GOPATH/src/bradford-hamilton/grpc-lru-cache
COPY . .
# Build for linux 64 bit. Omit the symbol table, debug information and the DWARF table for smaller binary, declare static
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"' -o /go/bin/server .
# Expose our GRPC service
EXPOSE 21000
# Run the server binary
ENTRYPOINT ["/go/bin/server"]
