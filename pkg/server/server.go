package server

import (
	"context"
	"errors"
	"log"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/mem"
	pb "github.com/bradford-hamilton/grpc-lru-cache/proto/cache"
)

// ErrEmptyCache is the default error message when asking for the MRU item or LRU item
var ErrEmptyCache = errors.New("error: cannot retrieve item - cache is empty")

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

// GetKeys retrieves all the available keys from cache
func (c *CacheServer) GetKeys(context.Context, *pb.Empty) (*pb.KeysRes, error) {
	keys := c.cache.Keys()
	strKeys := make([]string, len(keys))
	for i := range keys {
		strKeys[i] = keys[i].(string)
	}
	return &pb.KeysRes{Keys: strKeys}, nil
}

// Flush clears the cache and re-initializes it for use
func (c *CacheServer) Flush(context.Context, *pb.Empty) (*pb.Empty, error) {
	c.cache.Flush()
	return &pb.Empty{}, nil
}

// Cap returns the max number of items the cache can hold
func (c *CacheServer) Cap(context.Context, *pb.Empty) (*pb.CapRes, error) {
	return &pb.CapRes{Cap: int64(c.cache.Cap())}, nil
}

// Len returns the current number of items in the cache
func (c *CacheServer) Len(context.Context, *pb.Empty) (*pb.LenRes, error) {
	return &pb.LenRes{Len: int64(c.cache.Len())}, nil
}

// GetFirst gets the Most Recently Used item and if there are no items in the cache, returns an error
func (c *CacheServer) GetFirst(context.Context, *pb.Empty) (*pb.GetFirstOrLastRes, error) {
	val := c.cache.GetFront()
	if val == nil {
		return &pb.GetFirstOrLastRes{}, ErrEmptyCache
	}
	return &pb.GetFirstOrLastRes{Value: val.(string)}, nil
}

// GetLast gets the Least Recently Used item and if there are no items in the cache, returns an error
func (c *CacheServer) GetLast(context.Context, *pb.Empty) (*pb.GetFirstOrLastRes, error) {
	val := c.cache.GetBack()
	if val == nil {
		return &pb.GetFirstOrLastRes{}, ErrEmptyCache
	}
	return &pb.GetFirstOrLastRes{Value: val.(string)}, nil
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
