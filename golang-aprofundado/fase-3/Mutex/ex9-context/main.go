package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func filho(ctx context.Context, nome string, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
	fmt.Println(nome, "recebeu cancelamento", ctx.Err())
}

func main() {
	pai, cancelePai := context.WithCancel(context.Background())
	f1, cancelaF1 := context.WithCancel(pai)
	f2, cancelaF2 := context.WithCancel(pai)
	defer cancelaF1()
	defer cancelaF2()

	var wg sync.WaitGroup
	wg.Add(2)
	go filho(f1, "filho1", &wg)
	go filho(f2, "filho2", &wg)

	time.Sleep(100 * time.Millisecond)
	cancelePai()
	wg.Wait()
	fmt.Println("pai:", pai.Err())
}
