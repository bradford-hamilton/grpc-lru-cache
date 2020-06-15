package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bradford-hamilton/grpc-lru-cache/internal/server"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
	"google.golang.org/grpc"
)

// TODO: eventually take these as args
const cacheSize = 1024
const port = 21000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterCacheServiceServer(srv, server.NewCacheServer(cacheSize))

	fmt.Printf("Listening on port %d\n", port)
	srv.Serve(lis)
}
