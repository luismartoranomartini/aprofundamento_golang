package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	quadrados := make([]int, len(numeros))

	var wg sync.WaitGroup
	for i, n := range numeros {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			quadrados[i] = n * n
		}()
	}
	wg.Wait()
	fmt.Println(quadrados)
}
