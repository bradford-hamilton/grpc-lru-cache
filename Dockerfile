FROM golang:alpine AS builder

WORKDIR $GOPATH/src/bradford-hamilton/grpc-lru-cache

COPY . .

# Build for linux 64 bit. Omit the symbol table, debug information and the DWARF table for smaller binary, declare static
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"' -o /go/bin/server .

# Fetch upx https://github.com/upx/upx
RUN apk add upx

# Use upx to pack the binary even smaller
RUN upx /go/bin/server

# From scratch for the final image
FROM scratch

# Copy our static executable
COPY --from=builder /go/bin/server /go/bin/server

# Port choice to play nicely with one click deploy to GCP button
EXPOSE 8080

# Run the server binary
ENTRYPOINT ["/go/bin/server"]
