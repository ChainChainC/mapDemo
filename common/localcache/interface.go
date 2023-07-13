package localcache

import "time"

// Cache 缓存接口
type Cache interface {
	Set(key string, value interface{})
	SetWithExpiration(key string, value interface{}, expDuration time.Duration)
	Get(key string) (value interface{}, ok bool)
	Delete(key string) error
}
