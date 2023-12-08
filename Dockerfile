FROM golang:1.21.5
WORKDIR $GOPATH/src/bradford-hamilton/grpc-lru-cache
COPY . .

# Dependencies
RUN go mod tidy

# Build for linux x86-64
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $GOPATH/bin/server cmd/srv/cachesrv.go

# Expose port for our GRPC service
EXPOSE 21000

# Run the server binary
ENTRYPOINT ["/go/bin/server"]
