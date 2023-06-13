package main

import (
	"context"
	"flag"
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

const (
	defaultCacheSize = 4096
	port             = "21000"
)

func main() {
	size := flag.Int("size", defaultCacheSize, "underlying cache size in bytes")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	srv, lis, cacheSrv := registerGrpcCacheService(*size)

	go func() {
		fmt.Printf("Listening on port %s\n", port)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("error: %v", err)
		}
	}()

	<-sigs

	if err := cacheSrv.SaveToDisk(context.Background()); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func registerGrpcCacheService(size int) (*grpc.Server, net.Listener, *server.CacheServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	cacheSrv, err := server.NewCacheServer(size)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterCacheServiceServer(srv, cacheSrv)

	return srv, lis, cacheSrv
}
