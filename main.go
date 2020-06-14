package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/server"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
	"google.golang.org/grpc"
)

const port = 21000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterCacheServiceServer(srv, server.NewCacheServer(1024))

	fmt.Printf("Listening on port %d\n", port)
	srv.Serve(lis)
}
