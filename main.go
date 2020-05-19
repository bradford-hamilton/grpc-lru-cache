package main

import (
	"fmt"

	"github.com/bradford-hamilton/LRU-cache/pkg/lru"
)

func main() {
	c, err := lru.NewCacheClient(10)
	if err != nil {
		fmt.Println(err)
	}

	// Set some keys (keys can be any hashable type)
	if ok := c.Set("someKey", "someValue"); !ok {
		fmt.Println("didnt add: ", ok)
	}
	if ok := c.Set("price", 350000); !ok {
		fmt.Println("didnt add: ", ok)
	}
	if ok := c.Set(struct{ name string }{"daaaavid"}, 5.5); !ok {
		fmt.Println("didnt add: ", ok)
	}

	// Get some keys and print them
	if item, ok := c.Get("someKey"); ok {
		fmt.Println(item)
	}
	if item, ok := c.Get("price"); ok {
		fmt.Println(item)
	}
	if item, ok := c.Get(struct{ name string }{"daaaavid"}); ok {
		fmt.Println(item)
	}
	if item, ok := c.Get("someKeyThatIsntThere"); ok {
		fmt.Println(item)
	}

	// Print a slice of all available keys
	allKeys := c.Keys()
	fmt.Println(allKeys)
}

// TODO
// Maybe check type passed into Get/Set and if it isn't a "hashable" type, return error
// Concurrent testing
// More features
// Add benchmarking
