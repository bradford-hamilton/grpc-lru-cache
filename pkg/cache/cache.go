package cache

import (
	"container/list"
	"errors"
	"sync"
)

// ErrMinCacheSize is returned when a caller tries to create a new LRU cache with a capacity of less than one
var ErrMinCacheSize = errors.New("please provide an LRU cache capacity greater than or equal to 1")

// Cache represents an LRU cache and methods attached represent the main public API.
// Cache can be used in concurrent processes, it is thread safe.
type Cache struct {
	lru *lru
	mu  sync.Mutex // mutex for concurrent access to the cache
}

// New creates a new CacheClient with a max size provider by the caller.
func New(capacity int) (*Cache, error) {
	if capacity < 1 {
		return nil, ErrMinCacheSize
	}
	c := new(Cache)
	c.lru = &lru{
		cap:   capacity,
		ll:    list.New(),
		items: make(map[interface{}]*list.Element),
	}
	return c, nil
}

// Get handles finding a value by key in the cache. If found, it returns the value
// as well as true, signifying the cache hit. If no key is found it returns nil
// and false, signifying the cache miss.
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	item, ok := c.lru.get(key)
	c.mu.Unlock()
	return item, ok
}

// Set handles upserting the key in cache. The return values can be ignored if you are not
// interested in whether an Item was evicted or what that Item was. It can not error. If
// an item is evicted, it returns a copy of the item, as well as true to signify that the
// eviction happened. If nothing is evicted, the return Item will be a zero-value and false
// is returned to signify no eviction occurred.
func (c *Cache) Set(key, value interface{}) (Item, bool) {
	c.mu.Lock()
	item, ok := c.lru.set(key, value)
	c.mu.Unlock()
	return item, ok
}

// Flush clears the cache and re-initializes it for use.
func (c *Cache) Flush() {
	c.mu.Lock()
	c.lru.flush()
	c.mu.Unlock()
}

// Keys returns a slice of all the current keys available in cache.
func (c *Cache) Keys() []interface{} {
	c.mu.Lock()
	k := c.lru.keys()
	c.mu.Unlock()
	return k
}

// Cap returns the max number of items the cache can hold
func (c *Cache) Cap() int {
	return c.lru.cap
}

// Len returns the current number of items in the cache
func (c *Cache) Len() int {
	c.mu.Lock()
	l := len(c.lru.items)
	c.mu.Unlock()
	return l
}
