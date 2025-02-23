package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	data     map[string]cacheEntry
	mutex    sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		data:     make(map[string]cacheEntry),
		mutex:    sync.RWMutex{},
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	data, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		fmt.Println("< DELETING CACHE ENTRIES >")
		c.mutex.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
}
