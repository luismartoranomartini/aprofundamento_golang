package main

import (
	"fmt"
	"sync"
)

func main() {
	var uma sync.Once
	for i := 1; i <= 3; i++ {
		uma.Do(func() { fmt.Println("rodou com i =", i) })
	}
	fmt.Println("fim")
}
