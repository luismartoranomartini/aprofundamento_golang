package main

import (
	"context"
	"fmt"
	"time"
)

func trabalhador(ctx context.Context, nome string) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println(nome, "parando", ctx.Err())
			return
		case <-time.After(100 * time.Millisecond):
			n++
			fmt.Println(nome, "tique", n)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fim := make(chan struct{})
	go func() {
		trabalhador(ctx, "w1")
		close(fim)
	}()
	time.Sleep(350 * time.Millisecond)
	cancel()
	<-fim
	fmt.Println("main: fim")
}
