package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var contador atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("contador:", contador.Load())
}
