package main

import (
	"fmt"
	"sync"
	"time"
)

func tarefa(n int) {
	time.Sleep(1 * time.Second)
	fmt.Println("tarefa", n, "pronta")
}

func main() {
	ini := time.Now()
	for i := 1; i <= 3; i++ {
		tarefa(i)
	}
	fmt.Printf("sequencial: %.1fs\n", time.Since(ini).Seconds())

	ini = time.Now()
	var wg sync.WaitGroup
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		tarefa(i)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("i =", i)
		}()
	}
	wg.Wait()
	fmt.Printf("concorrente: %.1fs\n", time.Since(ini).Seconds())
}
