package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/cache"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
)

// CacheServer ...
type CacheServer struct{}

// Get ...
func (c *CacheServer) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	var item cache.Item
	b := []byte("hey")
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &item)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return &pb.GetRes{Value: []byte("hey")}, nil
}
