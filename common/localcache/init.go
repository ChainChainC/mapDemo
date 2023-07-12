package localcache

import (
	"sync"
	"time"

	glru "github.com/hashicorp/golang-lru"
	"github.com/patrickmn/go-cache"
)

var pool sync.Pool

func init() {
	pool.New = func() any {
		return &ttlValue{}
	}
}

// for test
var (
	testGoCache *cache.Cache

	testLRUCache5 *glru.Cache // 容量5
	testLRUCache3 *glru.Cache // 容量3
)

func init() {
	testGoCache = cache.New(10*time.Second, 10*time.Second)
}
