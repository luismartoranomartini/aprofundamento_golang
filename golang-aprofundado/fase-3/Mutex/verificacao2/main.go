package main

import (
	"fmt"
	"sync"
)

func main() {
	var a, b, c sync.Once

	a.Do(func() { fmt.Println("A1") })
	a.Do(func() { fmt.Println("A2") })
	b.Do(func() { fmt.Println("B1") })
	b.Do(func() { fmt.Println("B2") })
	c.Do(func() { fmt.Println("C1") })
	c.Do(func() { fmt.Println("C2") })
}
