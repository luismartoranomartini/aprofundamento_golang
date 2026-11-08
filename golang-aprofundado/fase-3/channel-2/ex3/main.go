package main

import (
	"fmt"
	"time"
)

func lenta() <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "resultado"
	}()
	return ch
}

func main() {
	ini := time.Now()
	select {
	case r := <-lenta():
		fmt.Println("ok", r)
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("timeout depois de %.1fs\n", time.Since(ini).Seconds())
	}
}
