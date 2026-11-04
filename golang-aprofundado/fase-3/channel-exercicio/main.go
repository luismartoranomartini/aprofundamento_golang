package main

import (
	"fmt"
	"time"
)

func esperar(ch <-chan int, limite time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(limite):
		return 0, false
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 42
	}()

	fmt.Println(esperar(ch1, 500*time.Millisecond))
	fmt.Println(esperar(ch2, 500*time.Millisecond))
}
