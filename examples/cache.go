package examples

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/bradford-hamilton/grpc-lru-cache/pkg/mem"
)

func One() {
	c, err := mem.NewLRUCache(5000)
	if err != nil {
		fmt.Println(err)
	}

	// Set some keys (keys can be any hashable type)
	if _, ok := c.Set("someKey", "someValue"); !ok {
		fmt.Println("didnt add: ", ok)
	}
	if _, ok := c.Set("price", "350000"); !ok {
		fmt.Println("didnt add: ", ok)
	}

	// Get some keys and print them
	if item, ok := c.Get("someKey"); ok {
		fmt.Println(item)
		fmt.Println(ok)
	}
	if item, ok := c.Get("price"); ok {
		fmt.Println(item)
		fmt.Println(ok)
	}

	// Getting a key that does not exist
	if item, ok := c.Get("someKeyThatIsntThere"); ok {
		fmt.Println(item)
		fmt.Println(ok)
	}

	// Print a slice of all available keys
	fmt.Println(c.Keys())
}

func Two() {
	c, err := mem.NewLRUCache(5000)
	if err != nil {
		fmt.Println(err)
	}

	// Example of 10,000 go routines writing safely
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go setItem(c, i, &wg)
	}
	wg.Wait()

	// Example of 10,000 go routines reading safely
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go getItem(c, i, &wg)
	}
	wg.Wait()

	// Print all available keys after waiting for go routines to finish
	fmt.Println(c.Keys())
}

func setItem(cache *mem.LRUCache, i int, wg *sync.WaitGroup) {
	cache.Set("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	wg.Done()
}

func getItem(cache *mem.LRUCache, i int, wg *sync.WaitGroup) {
	cache.Get("key" + strconv.Itoa(i))
	wg.Done()
}

// TODO
// Maybe check type passed into Get/Set and if it isn't a "hashable" type, return error
// Concurrent testing
// More features
