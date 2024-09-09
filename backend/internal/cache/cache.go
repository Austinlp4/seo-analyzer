package cache

import (
	"sync"
	"time"

	"automated-seo-analyzer/backend/internal/models"
)

type Cache struct {
	sync.RWMutex
	items map[string]cacheItem
}

type cacheItem struct {
	value      *models.AnalysisResponse
	expiration time.Time
}

var (
	cache     *Cache
	cacheTTL  = 1 * time.Hour
	onceCache sync.Once
)

func GetCache() *Cache {
	onceCache.Do(func() {
		cache = &Cache{
			items: make(map[string]cacheItem),
		}
	})
	return cache
}

func (c *Cache) Set(key string, value *models.AnalysisResponse) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(cacheTTL),
	}
}

func (c *Cache) Get(key string) (*models.AnalysisResponse, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		delete(c.items, key)
		return nil, false
	}
	return item.value, true
}
