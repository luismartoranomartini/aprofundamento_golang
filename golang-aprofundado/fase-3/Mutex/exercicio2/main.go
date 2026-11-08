package main

import (
	"context"
	"fmt"
	"time"
)

func gerar(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1

	// go func() {
	// 	defer close(ch)
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		case <-time.After(100 * time.Millisecond):
	// 			ch <- n
	// 			n++
	// 		}
	// 	}
	//
	// }()
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
			}

			select {
			case ch <- n:
				n++
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()

	for n := range gerar(ctx) {
		fmt.Println(n)
	}
	fmt.Println("fim")
}
