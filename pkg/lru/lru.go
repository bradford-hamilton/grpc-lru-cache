package lru

import (
	"container/list"
	"errors"
	"sync"
)

// ErrMinCacheSize is returned when a caller tries to create a new LRU cache with a size of less than one
var ErrMinCacheSize = errors.New("please provide an LRU cache size greater than or equal to 1")

// CacheClient represents the interface that must be implemented in order to
type CacheClient interface {
	Set(key, value interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Keys() []interface{}
}

// Cache represents our LRU cache and implements the CacheClient interface
type cache struct {
	size  int
	list  *list.List
	items map[interface{}]*list.Element
	mu    sync.Mutex
}

// Item represents a single item from our LRU cache, which simply has a key and value
type Item struct {
	key   interface{}
	value interface{}
}

// NewCacheClient creates a new CacheClient with a max size based on the size arg passed by caller.
func NewCacheClient(size int) (CacheClient, error) {
	if size < 1 {
		return nil, ErrMinCacheSize
	}
	return &cache{
		size:  size,
		list:  list.New(),
		items: make(map[interface{}]*list.Element),
	}, nil
}

// Get handles finding the key in cache, moving it to the front of our linked list (making it
// the most recently used item), and returning it. If no key is found it returns nil and false
// which represents whether the query was "ok"
func (c *cache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.list.MoveToFront(el)
		if el.Value.(*Item).value == nil {
			return Item{}, true
		}
		return el.Value.(*Item).value, true
	}

	return nil, false
}

// Set handles upserting the key in cache.
// If the key is already present, move it to the front and make it most recent.
// If the key is not present, set the key, push it to the front of our list (make it most recent),
// and evicting the least recently used item if the list length is greater than the cache size.
func (c *cache) Set(key, value interface{}) bool {
	c.mu.Lock()
	if el, ok := c.items[key]; ok {
		c.mu.Unlock()
		c.list.MoveToFront(el)
		el.Value.(*Item).value = value
		return false
	}
	c.items[key] = c.list.PushFront(&Item{key, value})
	c.mu.Unlock()

	// Check and evict the least recently used item when appropriate
	if c.list.Len() > c.size {
		c.evictLRUItem()
	}
	return true
}

// Keys returns a slice of all the current keys available in cache.
func (c *cache) Keys() []interface{} {
	var i int
	keys := make([]interface{}, len(c.items))
	c.mu.Lock()
	for _, item := range c.items {
		keys[i] = item.Value.(*Item).key
		i++
	}
	c.mu.Unlock()
	return keys
}

// evictLRUItem looks for the last ("Back") item on our cache's linked list.
// If it is found, a call to evict that specific element from the list is made.
func (c *cache) evictLRUItem() {
	if el := c.list.Back(); el != nil {
		c.evictElement(el)
	}
}

// evictElement takes a ptr to a list element and removes it from the list.
// After removing it from the list, we remove it from our cache's items map.
func (c *cache) evictElement(el *list.Element) {
	c.list.Remove(el)
	item := el.Value.(*Item)

	// Keep critical sections as small as possible
	c.mu.Lock()
	delete(c.items, item.key)
	c.mu.Unlock()
}
