package mem

import (
	"container/list"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

const backupLocation = "/.grpc-lru-cache/data.csv"

type cache struct {
	cap   int                      // max number of items the cache can hold before needing to evict.
	ll    *list.List               // a doubly linked list.
	items map[string]*list.Element // map of keys -> doubly linked list elements
}

// Item represents a single item from our LRU cache, which simply has a key and value.
type Item struct {
	Key   string
	Value string
}

// set return values can be ignored if you are not concerned with
// whether an Item was evicted or what that Item was. It can not error.
func (c *cache) set(key, value string) (Item, bool) {
	// Check to see if the key is already in cache
	if el, ok := c.items[key]; ok {
		// Found: move the item to most recently used (front)
		// position in the list and set the new value for that key
		c.ll.MoveToFront(el)
		el.Value.(*Item).Value = value
		return Item{}, false
	}

	// Push a new Item to the front of the linked list and set
	// the returned element in the cache map
	c.items[key] = c.ll.PushFront(&Item{key, value})

	// Check if our cache is at capacity
	if c.ll.Len() == c.cap {
		// Evict the least recently used item (back of the list)
		// and return a copy of the evicted item to the caller
		c.evictElement(c.ll.Back())
		itm := c.ll.Back().Value.(*Item)
		return *itm, true
	}

	return Item{}, false
}

// get looks for the key in cache and returns it if found. The second
// return value (bool) tells the caller whether an Item was found or not.
func (c *cache) get(key string) (string, bool) {
	// Look for the key in cache
	if el, ok := c.items[key]; ok {
		// Cache hit: move the element to the front of the list and return
		// the value as well as true telling the caller it was found
		c.ll.MoveToFront(el)
		return el.Value.(*Item).Value, true
	}
	// Cache miss
	return "", false
}

// evictElement takes a ptr to a list element and removes it from the list.
// After removing it from the list, we remove it from our cache's items map.
func (c *cache) evictElement(el *list.Element) {
	c.ll.Remove(el)
	item := el.Value.(*Item)
	delete(c.items, item.Key)
}

// flush clears the lru's items map and re-initializes the lru's linked list
func (c *cache) flush() {
	for k := range c.items {
		delete(c.items, k)
	}
	c.ll.Init()
}

// keys returns all current available keys in the cache
func (c *cache) keys() []string {
	var i int
	keys := make([]string, len(c.items))
	for _, item := range c.items {
		keys[i] = item.Value.(*Item).Key
		i++
	}
	return keys
}

// getFront gets the Most Recently Used item, and if there
// are no items in the cache at all, it will return nil
func (c *cache) getFront() string {
	el := c.ll.Front()
	if el == nil {
		return ""
	}
	return el.Value.(*Item).Value
}

// getBack gets the Least Recently Used item, and if there are
// no items in the cache at all, it will return nil. It also
// moves the back item to the front because it's been accessed.
func (c *cache) getBack() string {
	el := c.ll.Back()
	if el == nil {
		return ""
	}
	// Ensure item gets moved to the front of the cache
	c.ll.MoveToFront(el)
	return el.Value.(*Item).Value
}

// writeToDisk handles find or create for the ~/.grpc-lru-cache directory
// and find or create for the data.csv file that resides inside. Finally
// it calls to write the CSV backup data file.
func (c *cache) writeToDisk() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	if err := createConfigDirIfNotExists(home); err != nil {
		return err
	}
	if err := createConfigFileIfNotExists(home); err != nil {
		return err
	}
	if err := c.writeCSVDataBackup(home); err != nil {
		return err
	}
	return nil
}

func (c *cache) writeCSVDataBackup(home string) error {
	f, err := os.OpenFile(home+backupLocation, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, val := range c.items {
		item := val.Value.(*Item)
		err := w.Write([]string{item.Key, item.Value})
		if err != nil {
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (c *cache) seedBackupDataIfAvailable() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	ok, err := userHasBackupData(home)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	csvfile, err := os.Open(home + backupLocation)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)

	// Iterate through the records reading each individual record from the CSV file until EOF.
	for {
		// Read individual record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) != 2 {
			return errors.New("backup data corrupted, please delete ~/.grpc-lru-cache/data.csv and have it regenerate")
		}
		c.set(record[0], record[1])
	}

	return nil
}

// userHasBackupData's only concern is returning whether or not there are any bytes written inside /.grpc-lru-cache/data.csv
func userHasBackupData(home string) (bool, error) {
	fInfo, err := os.Stat(home + backupLocation)
	if err != nil {
		return false, err
	}
	if fInfo.Size() == 0 {
		return false, err
	}
	return true, nil
}

func createConfigDirIfNotExists(home string) error {
	if _, err := os.Stat(home + "/.grpc-lru-cache"); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(home+"/.grpc-lru-cache", os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

func createConfigFileIfNotExists(home string) error {
	if _, err := os.Stat(home + backupLocation); err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(home + backupLocation)
			if err != nil {
				return err
			}
			f.Close()
			return nil
		}
	}
	return nil
}
