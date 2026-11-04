package main

import (
	"fmt"
	"sync"
)

func dizer(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("goroutine", n)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go dizer(i, &wg)
	}
	wg.Wait()
	fmt.Println("fim do main")
}
