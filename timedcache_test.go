package cache

import "testing"

func TestCacheWithExpiration(t *testing.T) {
	c := NewCacheWithExpiration[int]()
	c.StoreWithOptions("test", 5, WithExpiration(10))

	entry := c.(*expCache[int]).data["test"]
	// view exp data
	t.Log(entry.data, entry.expires.Format("2006 Jan 2 15:04:05"))

	c.Remove("test")

	x, found := c.Get("test")
	t.Log(x, found)
}
