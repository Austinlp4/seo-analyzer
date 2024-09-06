package cache

import (
	"sync"
	"time"

	"github.com/Austinlp4/seo-analyzer/backend/internal/models"
)

type Cache struct {
	items map[string]*cacheItem
	mu    sync.RWMutex
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
			items: make(map[string]*cacheItem),
		}
	})
	return cache
}

func (c *Cache) Set(key string, value *models.AnalysisResponse) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = &cacheItem{
		value:      value,
		expiration: time.Now().Add(cacheTTL),
	}
}

func (c *Cache) Get(key string) (*models.AnalysisResponse, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
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
