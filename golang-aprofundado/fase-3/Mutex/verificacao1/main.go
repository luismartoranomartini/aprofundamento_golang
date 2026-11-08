package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var once sync.Once
			once.Do(func() { fmt.Println("abrindo") })
		}()
	}
	wg.Wait()
}
