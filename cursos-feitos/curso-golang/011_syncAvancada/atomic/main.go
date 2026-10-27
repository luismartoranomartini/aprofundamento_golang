package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// mutex simplificado

func main() {
	var (
		flag int64
		wg   sync.WaitGroup
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Escrevendo...")
		time.Sleep(time.Second * 3)
		atomic.StoreInt64(&flag, 1)
		fmt.Println("Valor escrito com sucesso")
	}()

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&flag) == 0 {
			fmt.Println("Flag com valor 0")
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Println("Flag com valor 1")
	}()

	wg.Wait()

}
