package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bradford-hamilton/grpc-lru-cache/internal/server"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
	"google.golang.org/grpc"
)

// TODO: eventually take these as args
const cacheSize = 1024
const port = 21000

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	srv, lis := registerGrpcService()

	go func() {
		fmt.Printf("Listening on port %d\n", port)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("error: %v", err)
		}
	}()

	<-sigs
}

func registerGrpcService() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterCacheServiceServer(srv, server.NewCacheServer(cacheSize))
	return srv, lis
}
