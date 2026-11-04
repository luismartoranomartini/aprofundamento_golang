package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		a <- "A (100ms)"
	}()
	go func() {
		time.Sleep(300 * time.Millisecond)
		b <- "B (300ms)"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-a:
			fmt.Println("recebi de a:", m)
		case m := <-b:
			fmt.Println("recebi de b", m)
		}

	}
}
