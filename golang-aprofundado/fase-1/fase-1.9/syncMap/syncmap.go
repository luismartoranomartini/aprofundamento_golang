package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("a", 1)
	m.Store("b", 2)
	v, ok := m.Load("a")
	fmt.Println(v, ok)
	_, ok = m.Load("z")
	fmt.Println(ok)
	m.Delete("b")
	m.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
}
