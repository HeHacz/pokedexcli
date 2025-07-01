package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]CacheEntry),
		mu:    sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		if len(c.cache) == 0 {
			c.mu.Unlock()
			return
		}
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
