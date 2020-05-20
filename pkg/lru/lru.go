package lru

import (
	"container/list"
	"errors"
	"sync"
)

// ErrMinCacheSize is returned when a caller tries to create a new LRU cache with a capacity of less than one
var ErrMinCacheSize = errors.New("please provide an LRU cache capacity greater than or equal to 1")

// Cache represents an LRU cache and methods attached represent the main public API.
type Cache struct {
	cap   int                           // max number of items the cache can hold before needing to evict.
	ll    *list.List                    // a doubly linked list.
	mu    sync.Mutex                    // mutex for concurrent access to the cache
	items map[interface{}]*list.Element // items mapping of key ->
}

// Item represents a single item from our LRU cache, which simply has a key and value
type Item struct {
	key   interface{}
	value interface{}
}

// NewCache creates a new CacheClient with a max size based on the size arg passed by caller.
func NewCache(capacity int) (*Cache, error) {
	if capacity < 1 {
		return nil, ErrMinCacheSize
	}
	return &Cache{
		cap:   capacity,
		ll:    list.New(),
		items: make(map[interface{}]*list.Element),
	}, nil
}

// Get handles finding the key in cache, moving it to the front of our linked list (making it
// the most recently used item), and returning it. If no key is found it returns nil and false
// which represents whether the query was "ok"
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		if el.Value.(*Item).value == nil {
			return Item{}, false
		}
		return el.Value.(*Item).value, true
	}

	return nil, false
}

// Set handles upserting the key in cache.
// If the key is already present, move it to the front and make it most recent.
// If the key is not present, set the key, push it to the front of our list (make it most recent),
// and evicting the least recently used item if the list length is greater than the cache size.
func (c *Cache) Set(key, value interface{}) bool {
	c.mu.Lock()
	if el, ok := c.items[key]; ok {
		c.mu.Unlock()
		c.ll.MoveToFront(el)
		el.Value.(*Item).value = value
		return false
	}
	c.items[key] = c.ll.PushFront(&Item{key, value})
	c.mu.Unlock()

	// Check and evict the least recently used item when appropriate
	if c.ll.Len() > c.cap {
		c.evictLRUItem()
	}
	return true
}

// Flush handles clearing out the items map and re-initializing the cache's list
func (c *Cache) Flush() {
	for k := range c.items {
		delete(c.items, k)
	}
	c.ll.Init()
}

// Keys returns a slice of all the current keys available in cache.
func (c *Cache) Keys() []interface{} {
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

// Cap returns the max number of items the cache can hold
func (c *Cache) Cap() int {
	return c.cap
}

// Len returns the current number of items in the cache
func (c *Cache) Len() int {
	return len(c.items)
}

// evictLRUItem looks for the last ("Back") item on our cache's linked list.
// If it is found, a call to evict that specific element from the list is made.
func (c *Cache) evictLRUItem() {
	if el := c.ll.Back(); el != nil {
		c.evictElement(el)
	}
}

// evictElement takes a ptr to a list element and removes it from the list.
// After removing it from the list, we remove it from our cache's items map.
func (c *Cache) evictElement(el *list.Element) {
	c.ll.Remove(el)
	item := el.Value.(*Item)

	// Keep critical sections as small as possible
	c.mu.Lock()
	delete(c.items, item.key)
	c.mu.Unlock()
}
