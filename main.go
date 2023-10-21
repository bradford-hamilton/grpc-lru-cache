package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bradford-hamilton/grpc-lru-cache/internal/server"
	pb "github.com/bradford-hamilton/grpc-lru-cache/protos/cache-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterCacheServiceHandlerFromEndpoint(ctx, mux, "0.0.0.0:3000", opts)
	if err != nil {
		log.Fatalf("failed to register cache service http server: %+v", err)
	}

	go func() {
		fmt.Printf("Listening on port %s\n", "3000")
		if err := http.ListenAndServe(":3000", mux); err != nil {
			log.Fatalf("error: %v", err)
		}
	}()

	<-sigs

	if err := cacheSrv.SaveToDisk(ctx); err != nil {
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
