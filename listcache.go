package cache

type ListCache[T any] struct {
	data []T
}

func NewListCache[T any]() *ListCache[T] {
	c := ListCache[T]{}
	return &c
}

func (c *ListCache[T]) Get() []T {
	return c.data
}

func (c *ListCache[T]) Store(list []T) {
	c.data = list
}

func (c *ListCache[T]) Clear() {
	c.data = nil
}
