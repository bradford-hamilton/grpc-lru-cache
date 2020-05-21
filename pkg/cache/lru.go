package cache

import (
	"container/list"
)

type lru struct {
	cap   int                           // max number of items the cache can hold before needing to evict.
	ll    *list.List                    // a doubly linked list.
	items map[interface{}]*list.Element // map of keys -> doubly linked list elements
}

// Item represents a single item from our LRU cache, which simply has a key and value.
type Item struct {
	key   interface{}
	value interface{}
}

// set return values can be ignored if you are not concerned with
// whether an Item was evicted or what that Item was. It can not error.
func (lru *lru) set(key, value interface{}) (Item, bool) {
	// Check to see if the key is already in cache
	if el, ok := lru.items[key]; ok {
		// Found: move the item to most recently used (front)
		// position in the list and set the new value for that key
		lru.ll.MoveToFront(el)
		el.Value.(*Item).value = value
		return Item{}, false
	}

	// Push a new Item to the front of the linked list and set
	// the returned element in the cache map
	lru.items[key] = lru.ll.PushFront(&Item{key, value})

	// Check if our cache is at capacity
	if lru.ll.Len() == lru.cap {
		// Evict the least recently used item (back of the list)
		// and return a copy of the evicted item to the caller
		lru.evictElement(lru.ll.Back())
		itm := lru.ll.Back().Value.(*Item)
		return *itm, true
	}

	return Item{}, false
}

// get looks for the key in cache and returns it if found. The second
// return value (bool) tells the caller whether an Item was found or not.
func (lru *lru) get(key interface{}) (interface{}, bool) {
	// Look for the key in cache
	if el, ok := lru.items[key]; ok {
		// Cache hit: move the element to the front of the list and return
		// the value as well as true telling the caller it was found
		lru.ll.MoveToFront(el)
		return el.Value.(*Item).value, true
	}
	// Cache miss
	return nil, false
}

// evictElement takes a ptr to a list element and removes it from the list.
// After removing it from the list, we remove it from our cache's items map.
func (lru *lru) evictElement(el *list.Element) {
	lru.ll.Remove(el)
	item := el.Value.(*Item)
	delete(lru.items, item.key)
}

// flush clears the lru's items map and re-initializes the lru's linked list
func (lru *lru) flush() {
	for k := range lru.items {
		delete(lru.items, k)
	}
	lru.ll.Init()
}

// keys returns all current available keys in the cache
func (lru *lru) keys() []interface{} {
	var i int
	keys := make([]interface{}, len(lru.items))
	for _, item := range lru.items {
		keys[i] = item.Value.(*Item).key
		i++
	}
	return keys
}
