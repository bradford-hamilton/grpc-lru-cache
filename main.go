package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/bradford-hamilton/lru-cache/pkg/cache"
)

func main() {
	c, err := cache.New(1000)
	if err != nil {
		fmt.Println(err)
	}

	// Set some keys (keys can be any hashable type)
	if _, ok := c.Set("someKey", "someValue"); !ok {
		fmt.Println("didnt add: ", ok)
	}
	if _, ok := c.Set("price", 350000); !ok {
		fmt.Println("didnt add: ", ok)
	}
	if _, ok := c.Set(struct{ name string }{"daaaavid"}, 5.5); !ok {
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

	// Print a slice of all available keys (a few from above)
	fmt.Println(c.Keys())

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go setItem(c, i, &wg)
	}
	wg.Wait()

	// Print all available keys after waiting for go routines to finish
	fmt.Println(c.Keys())
}

func setItem(cache *cache.Cache, i int, wg *sync.WaitGroup) {
	cache.Set("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	wg.Done()
}

// TODO
// Maybe check type passed into Get/Set and if it isn't a "hashable" type, return error
// Concurrent testing
// More features
