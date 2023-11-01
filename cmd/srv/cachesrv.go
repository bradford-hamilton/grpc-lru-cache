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
	pb "github.com/bradford-hamilton/grpc-lru-cache/protos/cache-service"
	"google.golang.org/grpc"
)

const (
	defaultCacheSize = 1024 * 1024 // 1048576 - 8MB
	port             = "21000"
)

func main() {
	size := flag.Int("size", defaultCacheSize, "underlying cache size in bytes")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	srv, lis, cacheSrv := registerGrpcCacheService(*size)

	go func() {
		log.Printf("Listening on port %s\n", port)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("error: %v\n", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	<-sigs

	srv.GracefulStop()

	if err := cacheSrv.SaveToDisk(ctx); err != nil {
		log.Printf("error: %v\n", err)
	}
}

func registerGrpcCacheService(size int) (*grpc.Server, net.Listener, *server.CacheServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()
	cacheSrv, err := server.NewCacheServer(size)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterCacheServiceServer(srv, cacheSrv)

	return srv, lis, cacheSrv
}
