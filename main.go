package main

import (
	"fmt"

	"github.com/bradford-hamilton/LRU-cache/pkg/lru"
)

func main() {
	cache, err := lru.NewCacheClient(10)
	if err != nil {
		fmt.Println(err)
	}

	if ok := cache.Set("someKey", "someValue"); ok == false {
		fmt.Println("didnt add: ", ok)
	}

	item, ok := cache.Get("someKey")
	if !ok {
		fmt.Println(err)
	}

	item2, ok := cache.Get("someKeyThatIsntThere")
	if !ok {
		fmt.Println(err)
	}

	fmt.Printf("item1: %+v", item)
	fmt.Printf("item2: %+v", item2)
}

// TODO
// Maybe check type passed into Get/Set and if it isn't a "hashable" type, return error
// Concurrent testing
// More features
// Add benchmarking
