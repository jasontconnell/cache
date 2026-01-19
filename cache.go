package cache

import "strings"

type Cache[T any] struct {
	data map[string]T
}

func NewCache[T any]() *Cache[T] {
	c := Cache[T]{}
	c.data = make(map[string]T)
	return &c
}

func (c *Cache[T]) Get(key string) (T, bool) {
	var r T
	found := false
	if t, ok := c.data[key]; ok {
		r = t
		found = true
	}
	return r, found
}

func (c *Cache[T]) Store(key string, obj T) bool {
	c.data[key] = obj
	return true
}

func (c *Cache[T]) Remove(key string) {
	delete(c.data, key)
}

func (c *Cache[T]) ClearLike(keyLike string) {
	for k := range c.data {
		if strings.Contains(k, keyLike) {
			c.Remove(k)
		}
	}
}

func (c *Cache[T]) Clear() {
	c.data = make(map[string]T)
}
