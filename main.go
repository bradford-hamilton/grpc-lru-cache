package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/server"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 21000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cs := server.CacheServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterCacheServiceServer(grpcServer, &cs)

	grpcServer.Serve(lis)
}
