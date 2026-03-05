/* * This file defines a simple in-memory cache with TTL (Time To Live) functionality.
 * It allows storing key-value pairs with an expiration time, and provides methods
 * to set, get, delete, and clear cache entries. Additionally, it includes a cleanup
 * mechanism to remove expired entries at regular intervals.
 */

package storage

import (
	"sync"
	"time"
)

/**
 * CacheEntry represents a single entry in the cache, containing the value and its expiration time.
 */


type CacheEntry struct {
	value interface{}
	expiration time.Time
}

/**
 * Cache is a thread-safe in-memory cache with TTL functionality.
 * It uses a map to store cache entries and a RWMutex for concurrency control.
 */

type Cache struct {
	entries map[string]*CacheEntry
	mu      sync.RWMutex
	ttl     time.Duration
}

/**
 * NewCache creates a new Cache instance with the specified TTL (Time To Live) for cache entries.
 * It initializes the entries map and starts the cleanup process to remove expired entries.
 */

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]*CacheEntry),
		ttl:     ttl,
	}
	cache.StartCleanup(ttl)
	return cache
}


/**
 * Set adds a new entry to the cache with the specified key and value.
 * The entry will expire after the duration defined by the Cache's TTL.
 */

func (c *Cache) Set(key string, valeur interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = &CacheEntry{
		value:      valeur,
		expiration: time.Now().Add(c.ttl),
	}
}


/**
 * Get retrieves the value associated with the specified key from the cache.
 * It returns the value and a boolean indicating whether the key was found and is not expired.
 */

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, found := c.entries[key]
	if !found || time.Now().After(entry.expiration) {
		return nil, false
	}
	return entry.value, true
}


/**
 * Delete removes the entry associated with the specified key from the cache.
 */

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, key)
}


/** * Clear removes all entries from the cache.
 */

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]*CacheEntry)
}


/** * StartCleanup starts a background goroutine that periodically checks for expired entries
 * and removes them from the cache. The cleanup interval is defined by the provided duration.
 */

func (c *Cache) StartCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Now().After(entry.expiration) {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}