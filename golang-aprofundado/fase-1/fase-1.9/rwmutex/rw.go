package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu sync.RWMutex
	d  map[string]int
}

func (c *Cache) Get(k string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.d[k]
	return v, ok
}

func (c *Cache) Set(k string, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.d[k] = v
}

func main() {
	c := &Cache{d: map[string]int{}}
	c.Set("a", 1)
	fmt.Println(c.Get("a"))
	fmt.Println(c.Get("z"))
}
