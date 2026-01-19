package cache

type ListCache[T any] interface {
	Get() []T
	Store(list []T)
	Clear()
}

type basicListCache[T any] struct {
	data []T
}

func NewListCache[T any]() ListCache[T] {
	c := basicListCache[T]{}
	return &c
}

func (c *basicListCache[T]) Get() []T {
	return c.data
}

func (c *basicListCache[T]) Store(list []T) {
	c.data = list
}

func (c *basicListCache[T]) Clear() {
	c.data = nil
}
