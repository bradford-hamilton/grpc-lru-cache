package server

import (
	"context"
	"log"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/mem"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
)

// NewCacheServer creates a new *LRUCache with a caller-provided size, attaches it to a new
// CacheServer, and returns it to the caller. This allows us to use the LRUCache in only a
// string -> string capacity, as it normally accpets any type for keys and values.
func NewCacheServer(size int) *CacheServer {
	lru, err := mem.NewLRUCache(size)
	if err != nil {
		log.Panicf("failed to create new cache, err: %v", err)
	}

	return &CacheServer{cache: lru}
}

// CacheServer implements our CacheService grpc server. It contains an unexported cache
// field which get's hydrated with a new LRUCache when you call 'NewCacheServer'.
type CacheServer struct {
	cache *mem.LRUCache
}

// Get looks for an item in cache by its key and returns a "CacheHit" attribute for the caller to check against.
func (c *CacheServer) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	val, ok := c.cache.Get(req.Key)
	if !ok {
		return &pb.GetRes{CacheHit: false}, nil
	}
	return &pb.GetRes{Value: val.(string), CacheHit: true}, nil
}
