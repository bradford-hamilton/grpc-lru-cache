package mem

import (
	"container/list"
	"errors"
	"sync"
)

// ErrMinCacheSize is returned when a caller tries to create a new LRU cache with a capacity of less than one
var ErrMinCacheSize = errors.New("please provide an LRU cache capacity greater than or equal to 1")

// LRUCache represents an LRU cache and methods attached represent the main public API.
// LRUCache can be used in concurrent processes, it is thread safe.
type LRUCache struct {
	cache *cache
	mu    sync.Mutex // mutex for concurrent access to the cache
}

// NewLRUCache creates a new LRUCache with a max size provider by the caller.
func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity < 1 {
		return nil, ErrMinCacheSize
	}
	l := new(LRUCache)
	l.cache = &cache{
		cap:   capacity,
		ll:    list.New(),
		items: make(map[interface{}]*list.Element),
	}
	return l, nil
}

// Get handles finding a value by key in the cache. If found, it returns the value
// as well as true, signifying the cache hit. If no key is found it returns nil
// and false, signifying the cache miss.
func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	l.mu.Lock()
	item, ok := l.cache.get(key)
	l.mu.Unlock()
	return item, ok
}

// Set handles upserting the key in cache. The return values can be ignored if you are not
// interested in whether an Item was evicted or what that Item was. It can not error. If
// an item is evicted, it returns a copy of the item, as well as true to signify that the
// eviction happened. If nothing is evicted, the return Item will be a zero-value and false
// is returned to signify no eviction occurred.
func (l *LRUCache) Set(key, value interface{}) (Item, bool) {
	l.mu.Lock()
	item, ok := l.cache.set(key, value)
	l.mu.Unlock()
	return item, ok
}

// Flush clears the cache and re-initializes it for use.
func (l *LRUCache) Flush() {
	l.mu.Lock()
	l.cache.flush()
	l.mu.Unlock()
}

// Keys returns a slice of all the current keys available in cache.
func (l *LRUCache) Keys() []interface{} {
	l.mu.Lock()
	k := l.cache.keys()
	l.mu.Unlock()
	return k
}

// Cap returns the max number of items the cache can hold
func (l *LRUCache) Cap() int {
	return l.cache.cap
}

// Len returns the current number of items in the cache
func (l *LRUCache) Len() (length int) {
	l.mu.Lock()
	length = len(l.cache.items)
	l.mu.Unlock()
	return
}

// GetFront gets the Most Recently Used item, and if there
// are no items in the cache at all, it will return nil
func (l *LRUCache) GetFront() interface{} {
	l.mu.Lock()
	item := l.cache.getFront()
	l.mu.Unlock()
	return item
}

// GetBack gets the Least Recently Used item, and if there
// are no items in the cache at all, it will return nil
func (l *LRUCache) GetBack() interface{} {
	l.mu.Lock()
	item := l.cache.getBack()
	l.mu.Unlock()
	return item
}
