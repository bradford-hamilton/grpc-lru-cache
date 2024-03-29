package server

import (
	"context"
	"errors"
	"log"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/mem"
	pb "github.com/bradford-hamilton/grpc-lru-cache/protos/cache-service"
)

// ErrEmptyCache is the default error message when asking for the MRU item or LRU item.
var ErrEmptyCache = errors.New("error: cannot retrieve item - cache is empty")

// NewCacheServer creates a new *LRUCache with a caller-provided cap, attaches it to a new
// CacheServer, and returns it to the caller. This allows us to use the LRUCache in only a
// string -> string capacity, as it normally accpets any type for keys and values.
func NewCacheServer(capacity int) (*CacheServer, error) {
	lru, err := mem.NewLRUCache(capacity)
	if err != nil {
		log.Panicf("failed to create new cache, err: %v", err)
	}
	if err := lru.SeedBackupDataIfAvailable(); err != nil {
		return nil, err
	}
	return &CacheServer{cache: lru}, nil
}

// CacheServer implements our CacheService grpc server. It contains an unexported cache
// field which get's hydrated with a new LRUCache when you call 'NewCacheServer'.
type CacheServer struct {
	cache *mem.LRUCache
}

// Get looks for an item in cache by its key and returns a "CacheHit" attribute for the caller to check against.
func (c *CacheServer) Get(_ context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	val, ok := c.cache.Get(req.GetKey())
	if !ok {
		return &pb.GetRes{CacheHit: false}, nil
	}
	return &pb.GetRes{Value: val, CacheHit: true}, nil
}

// Set adds the provided item in cache. If an item was evicted because of this call, it will return
// Evicted == true as well as the the evicted item. If you are only concerned with setting the item
// in cache and you don't care whether anything was evicted, feel free to ignore both return vars.
func (c *CacheServer) Set(_ context.Context, item *pb.Item) (*pb.SetRes, error) {
	evictedItem, evicted := c.cache.Set(item.GetKey(), item.GetValue())
	if evicted {
		return evictionRes(evictedItem), nil
	}
	return &pb.SetRes{}, nil
}

// GetKeys retrieves all the available keys from cache.
func (c *CacheServer) GetKeys(_ context.Context, _ *pb.Empty) (*pb.KeysRes, error) {
	return &pb.KeysRes{Keys: c.cache.Keys()}, nil
}

// GetFirst gets the Most Recently Used item and if there are no items in the cache, returns an error.
func (c *CacheServer) GetFirst(_ context.Context, _ *pb.Empty) (*pb.GetFirstOrLastRes, error) {
	val := c.cache.GetFront()
	if val == "" {
		return &pb.GetFirstOrLastRes{}, ErrEmptyCache
	}
	return &pb.GetFirstOrLastRes{Value: val}, nil
}

// GetLast gets the Least Recently Used item and if there are no items in the cache, returns an error.
func (c *CacheServer) GetLast(_ context.Context, _ *pb.Empty) (*pb.GetFirstOrLastRes, error) {
	val := c.cache.GetBack()
	if val == "" {
		return &pb.GetFirstOrLastRes{}, ErrEmptyCache
	}
	return &pb.GetFirstOrLastRes{Value: val}, nil
}

// Flush clears the cache and re-initializes it for use.
func (c *CacheServer) Flush(_ context.Context, _ *pb.Empty) (*pb.Empty, error) {
	c.cache.Flush()
	return &pb.Empty{}, nil
}

// SaveToDisk will write all current key pairs to a CSV file in ~/.grpc-lru-cache/data.csv.
func (c *CacheServer) SaveToDisk(_ context.Context) error {
	if err := c.cache.SaveToDisk(); err != nil {
		return err
	}
	return nil
}

// Cap returns the max number of items the cache can hold.
func (c *CacheServer) Cap(_ context.Context, _ *pb.Empty) (*pb.CapRes, error) {
	return &pb.CapRes{Cap: int64(c.cache.Cap())}, nil
}

// Len returns the current number of items in the cache.
func (c *CacheServer) Len(_ context.Context, _ *pb.Empty) (*pb.LenRes, error) {
	return &pb.LenRes{Len: int64(c.cache.Len())}, nil
}

func evictionRes(evicted mem.Item) *pb.SetRes {
	return &pb.SetRes{
		EvictedItem: &pb.Item{
			Key:   evicted.Key,
			Value: evicted.Value,
		},
		Evicted: true,
	}
}
