package main

import (
	"context"
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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "21000"
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	srv, lis, cacheSrv := registerGrpcCacheService(port)

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

func registerGrpcCacheService(port string) (*grpc.Server, net.Listener, *server.CacheServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	cacheSrv, err := server.NewCacheServer(cacheSize)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterCacheServiceServer(srv, cacheSrv)

	return srv, lis, cacheSrv
}
