package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var c atomic.Int64

	fmt.Println("Load inicial:", c.Load())

	novo := c.Add(1)
	fmt.Println("Add(1) devolveu:", novo)

	novo = c.Add(5)
	fmt.Println("Add(5) devolveu:", novo)

	fmt.Println("Load agora:", c.Load())

	c.Store(100)
	fmt.Println("Depois do Store(100):", c.Load())

	ok := c.CompareAndSwap(100, 200)
	fmt.Println("CompareAndSwap(100, 200):", ok, "->", c.Load())

	ok = c.CompareAndSwap(100, 300)
	fmt.Println("CompareAndSwap(100, 300)", ok, "->", c.Load())
}
