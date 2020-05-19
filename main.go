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

	// Set some keys
	if ok := c.Set("someKey", "someValue"); ok == false {
		fmt.Println("didnt add: ", ok)
	}
	if ok := c.Set("price", 350000); ok == false {
		fmt.Println("didnt add: ", ok)
	}
	if ok := c.Set(struct{ name string }{"daaaavid"}, 5.5); ok == false {
		fmt.Println("didnt add: ", ok)
	}

	// Get some keys
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

	allKeys := c.Keys()
	fmt.Println(allKeys)
}

// TODO
// Maybe check type passed into Get/Set and if it isn't a "hashable" type, return error
// Concurrent testing
// More features
// Add benchmarking
