package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				m[j] = j
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
