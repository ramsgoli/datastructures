package lrucache

import (
	"errors"
	"fmt"
	"github.com/ramsgoli/datastructures/v2/linkedlist"
)

type LRUCache struct {
	size       int
	cache      *map[string]string
	cachedKeys *linkedlist.LinkedList
}

func New(size int) *LRUCache {
	return &LRUCache{size, &map[string]string{}, linkedlist.New()}
}

// Get get's the value of the specified key
func (c *LRUCache) Get(key string) (string, error) {
	if value, ok := (*c.cache)[key]; ok {
		// remove the key from the cached keys
		c.cachedKeys.Remove(key)

		// and bump it to the back
		c.cachedKeys.Insert(key)

		return value, nil
	}

	return "", errors.New("Key not found")
}

// Set set's the specified key and value in the cache
func (c *LRUCache) Set(key string, value string) {
	(*c.cache)[key] = value
	c.cachedKeys.Insert(key)

	// if cache size exceeds n, remove lease recently used
	if c.cachedKeys.Size > c.size {
		lru := c.cachedKeys.Head.Key
		delete(*c.cache, lru.(string))
		c.cachedKeys.Remove(lru)
	}
}

func (c *LRUCache) Print() {
	fmt.Printf("Number of keys in cache: %d\n", c.cachedKeys.Size)
	for key, value := range *c.cache {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
