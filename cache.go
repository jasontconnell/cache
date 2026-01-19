package cache

import (
	"strings"
)

const (
	NoExpiration int = -1
)

type Cache[T any] interface {
	Get(key string) (T, bool)
	Store(key string, obj T) bool
	StoreWithOptions(key string, obj T, opt ...Option) bool
	Remove(key string)
	ClearLike(keyLike string)
	Clear()
}

type basicCache[T any] struct {
	data map[string]T
}

func NewCache[T any]() Cache[T] {
	c := basicCache[T]{}
	c.data = make(map[string]T)
	return &c
}

func (c *basicCache[T]) Get(key string) (T, bool) {
	var r T
	found := false
	if t, ok := c.data[key]; ok {
		r = t
		found = true
	}
	return r, found
}

func (c *basicCache[T]) Store(key string, obj T) bool {
	c.data[key] = obj
	return true
}

func (c *basicCache[T]) StoreWithOptions(key string, obj T, opts ...Option) bool {
	return c.Store(key, obj)
}

func (c *basicCache[T]) Remove(key string) {
	delete(c.data, key)
}

func (c *basicCache[T]) ClearLike(keyLike string) {
	for k := range c.data {
		if strings.Contains(k, keyLike) {
			c.Remove(k)
		}
	}
}

func (c *basicCache[T]) Clear() {
	c.data = make(map[string]T)
}
