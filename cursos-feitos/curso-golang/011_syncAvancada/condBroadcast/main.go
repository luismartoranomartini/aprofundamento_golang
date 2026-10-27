package main

import (
	"fmt"
	"sync"
	"time"
)

type Trava struct {
	liberada bool
	mu       sync.Mutex
	cond     *sync.Cond
}

func NovaTrava() *Trava {
	t := &Trava{}
	t.cond = sync.NewCond(&t.mu)
	return t
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%dms] %s\n", time.Since(start).Milliseconds(), msg)
}

func worker(id int, t *Trava, wg *sync.WaitGroup) {
	defer wg.Done()
	t.mu.Lock()
	// defer t.mu.Unlock()

	for !t.liberada {
		logf(fmt.Sprintf("Worker %d aguardando a trava liberar...", id))
		t.cond.Wait()
	}
	t.mu.Unlock()
	logf(fmt.Sprintf("Worker %d executando...", id))
	time.Sleep(time.Millisecond * 100)
}

func main() {
	{ // bloco de escopo
		logf("===Broadcast===")
		var (
			qtdWorker = 4
			trava     = NovaTrava()
			wg        sync.WaitGroup
		)
		wg.Add(qtdWorker)

		for i := 1; i <= qtdWorker; i++ {
			go worker(i, trava, &wg)
		}
		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("trava liberada")
		trava.cond.Broadcast()
		trava.mu.Unlock()
		wg.Wait()
	}
	{ // bloco de escopo
		logf("=== Signal ===")
		var (
			qtdWorker = 4
			trava     = NovaTrava()
			wg        sync.WaitGroup
		)
		wg.Add(qtdWorker)

		for i := 1; i <= qtdWorker; i++ {
			go worker(i, trava, &wg)
		}
		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("trava liberada")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("trava liberada")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("trava liberada")
		trava.cond.Signal()
		trava.mu.Unlock()

		time.Sleep(time.Millisecond * 500)
		trava.mu.Lock()
		trava.liberada = true
		logf("trava liberada")
		trava.cond.Signal()
		trava.mu.Unlock()
		wg.Wait()
	}

}
