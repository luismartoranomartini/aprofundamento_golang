package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	quadrados := make([]int, len(numeros))

	var g errgroup.Group
	for i, n := range numeros {
		g.Go(func() error {
			quadrados[i] = n * n
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println(quadrados)
}
