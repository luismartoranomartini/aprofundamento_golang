package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Resultado struct {
	Job    int
	Worker int
	Valor  int
}

func trabalhador(id int, jobs <-chan int, resultados chan<- Resultado, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(50 * time.Millisecond)
		resultados <- Resultado{Job: j, Worker: id, Valor: j * j}
	}
}

func main() {
	jobs := make(chan int)
	resultados := make(chan Resultado)

	var wg sync.WaitGroup
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go trabalhador(w, jobs, resultados, &wg)
	}
	go func() {
		for j := 1; j <= 12; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(resultados)
	}()
	var todos []Resultado
	for r := range resultados {
		todos = append(todos, r)
	}
	sort.Slice(todos, func(a, b int) bool { return todos[a].Job < todos[b].Job })
	usados := map[int]bool{}
	for _, r := range todos {
		usados[r.Worker] = true
		fmt.Printf("job %2d -> %3d(worker %d)\n", r.Job, r.Valor, r.Worker)
	}
	fmt.Println("workers que trabalharam:", len(usados))
}
