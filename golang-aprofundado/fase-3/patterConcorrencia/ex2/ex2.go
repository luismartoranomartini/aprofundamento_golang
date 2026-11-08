package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func trabalhador(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case <-time.After(100 * time.Millisecond):
				_ = j
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	jobs := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go trabalhador(ctx, jobs, &wg)
	}
	enviadas := 0
produtor:
	for j := 1; j <= 1000; j++ {
		select {
		case jobs <- j:
			enviadas++
		case <-ctx.Done():
			break produtor
		}
	}
	close(jobs)
	wg.Wait()

	time.Sleep(50 * time.Millisecond)
	fmt.Println("tarefas aceitar antes do cancelamento:", enviadas)
	fmt.Println("goroutines vivas:", runtime.NumGoroutine())
}
