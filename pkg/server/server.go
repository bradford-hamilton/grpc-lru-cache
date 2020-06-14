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

// Set adds the provided item in cache. If an item was evicted because of this call, it will return
// Evicted == true as well as the the evicted item. If you are only concerned with setting the item
// in cache and you don't care whether anything was evicted, feel free to ignore both return vars.
func (c *CacheServer) Set(ctx context.Context, item *pb.Item) (*pb.SetRes, error) {
	evictedItem, evicted := c.cache.Set(item.Key, item.Value)
	if evicted {
		return evictionRes(evictedItem), nil
	}
	return &pb.SetRes{}, nil
}

func evictionRes(evicted mem.Item) *pb.SetRes {
	return &pb.SetRes{
		EvictedItem: &pb.Item{
			Key:   evicted.Key.(string),
			Value: evicted.Value.(string),
		},
		Evicted: true,
	}
}
