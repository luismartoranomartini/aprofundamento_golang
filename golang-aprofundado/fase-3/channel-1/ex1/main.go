package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Olá pelo canal"
	}()

	msg := <-ch
	fmt.Println(msg)
}
