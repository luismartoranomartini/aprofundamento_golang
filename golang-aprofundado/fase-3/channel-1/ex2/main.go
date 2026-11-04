package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ini := time.Now()

	go func() {
		fmt.Println("enviando...")
		ch <- 42
		fmt.Printf("enviado, depos de %.0fs\n", time.Since(ini).Seconds())
	}()
	time.Sleep(1 * time.Second)
	v := <-ch
	fmt.Println("recebido:", v)
	time.Sleep(100 * time.Millisecond)
}
