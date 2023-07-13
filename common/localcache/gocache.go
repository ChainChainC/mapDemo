package localcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// goCache goCache缓存
type goCache struct {
	cache *cache.Cache
}

// NewGoCache new cache.
func NewGoCache(defaultExpiration, cleanupInterval time.Duration) (Cache, error) {
	return &goCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}, nil
}

// Set 设置缓存
func (c *goCache) Set(key string, value interface{}) {
	c.cache.Set(key, value, 0)
}

// SetWithExpiration 给key设置过期时间
func (c *goCache) SetWithExpiration(key string, value interface{}, expDuration time.Duration) {
	c.cache.Set(key, value, expDuration)
}

// Get 获取缓存
func (c *goCache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

// Delete 删除
func (c *goCache) Delete(key string) error {
	c.cache.Delete(key)
	return nil
}
