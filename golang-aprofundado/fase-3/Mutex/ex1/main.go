package main

import (
	"fmt"
	"sync"
)

func main() {
	contador := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador++
		}()
	}
	wg.Wait()
	fmt.Println("contador:", contador)
}
