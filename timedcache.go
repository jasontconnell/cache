package cache

import (
	"strings"
	"sync"
	"time"
)

func NewCacheWithExpiration[T any]() Cache[T] {
	c := expCache[T]{}
	c.data = make(map[string]cacheEntry[T])
	return &c
}

type expCache[T any] struct {
	data map[string]cacheEntry[T]
	mu   sync.Mutex
}

type cacheEntry[T any] struct {
	data        T
	expires     time.Time
	nonExpiring bool
}

func (c *expCache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var r T
	found := false
	if t, ok := c.data[key]; ok {
		if t.nonExpiring || t.expires.After(time.Now()) {
			r = t.data
			found = true
		}
	}
	return r, found
}

func (c *expCache[T]) Store(key string, obj T) bool {
	return c.StoreWithOptions(key, obj, nil)
}

func (c *expCache[T]) StoreWithOptions(key string, obj T, opts ...Option) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	minutes := NoExpiration
	for _, opt := range opts {
		switch opt.(type) {
		case expOption:
			minutes = opt.GetValue().(int)
		}
	}

	nonExpiring := minutes == NoExpiration
	exptime := time.Time{}
	if !nonExpiring {
		exptime = time.Now().Add(time.Minute * time.Duration(minutes))
	}
	c.data[key] = cacheEntry[T]{
		data:        obj,
		nonExpiring: nonExpiring,
		expires:     exptime,
	}
	return true
}

func (c *expCache[T]) Remove(key string) {
	delete(c.data, key)
}

func (c *expCache[T]) ClearLike(keyLike string) {
	for k := range c.data {
		if strings.Contains(k, keyLike) {
			c.Remove(k)
		}
	}
}

func (c *expCache[T]) Clear() {
	c.data = make(map[string]cacheEntry[T])
}
