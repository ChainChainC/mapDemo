package localcache

import (
	"time"

	glru "github.com/hashicorp/golang-lru"
)

type ttlValue struct {
	V      interface{}
	expire *time.Time
}

func newValue(v interface{}) *ttlValue {
	tv := pool.Get().(*ttlValue)
	tv.V = v
	return tv
}

func release(tv *ttlValue) {
	tv.V = nil
	tv.expire = nil
	pool.Put(tv)
}

// NewLRUCache return LRU Cache
func NewLRUCache(capacity int64, d time.Duration) (Cache, error) {
	cache, err := glru.NewWithEvict(int(capacity), func(key interface{}, value interface{}) {
		if tv, ok := value.(*ttlValue); ok {
			release(tv)
		}
	})
	if err == nil {
		return &LRUCache{
			cache:   cache,
			timeout: d,
		}, nil
	}
	cache.Len()
	return nil, err
}

// LRUCache LRU缓存
type LRUCache struct {
	cache   *glru.Cache
	timeout time.Duration
}

// Set 写入
func (c *LRUCache) Set(k string, v interface{}) {
	tv := newValue(v)
	if c.timeout > 0 {
		expire := time.Now().Add(c.timeout)
		tv.expire = &expire
	}
	c.cache.Add(k, tv)
}

// SetWithExpiration 给key设置过期时间
func (c *LRUCache) SetWithExpiration(k string, v interface{}, expDuration time.Duration) {
	panic("implement me")
}

// Get 查询
func (c *LRUCache) Get(k string) (interface{}, bool) {
	v, ok := c.cache.Get(k)
	if ok {
		if tv, is := v.(*ttlValue); is {
			if tv.expire != nil && tv.expire.Before(time.Now()) {
				c.cache.Remove(k)
			} else {
				return tv.V, true
			}
		}
	}
	return nil, false
}

// Delete 删除
func (c *LRUCache) Delete(k string) error {
	c.cache.Remove(k)
	return nil
}

// Len 缓存长度
func (c *LRUCache) Len() int {
	return c.cache.Len()
}
