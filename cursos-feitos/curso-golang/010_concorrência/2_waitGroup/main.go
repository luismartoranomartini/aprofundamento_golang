package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		escrever("Olá, mundo")
		wg.Done() // -1
	}()

	go func() {
		escrever("Programar em Golang")
		wg.Done() // -1
	}()

	wg.Wait() //

}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
