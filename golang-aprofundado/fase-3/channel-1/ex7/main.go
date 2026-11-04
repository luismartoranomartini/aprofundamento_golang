package main

import (
	"fmt"
	"time"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	resultados := make(chan int)

	for _, n := range numeros {
		go func() {
			time.Sleep(500 * time.Millisecond)
			resultados <- n * n
		}()
	}
	for range numeros {
		fmt.Println(<-resultados)
	}
}
