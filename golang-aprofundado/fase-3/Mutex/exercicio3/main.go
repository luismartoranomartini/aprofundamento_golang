package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	resultado := make([]int, 4)

	for id := 1; id <= 4; id++ {
		g.Go(func() error {
			if id == 3 {
				return fmt.Errorf("A tarefa %d falhou", id)
			}
			resultado[id-1] = id * 10
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("erro:", err)
	}
	fmt.Println(resultado)
}
