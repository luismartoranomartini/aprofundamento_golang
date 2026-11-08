package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	g.SetLimit(2)

	var ativas, maximo atomic.Int64
	for i := 1; i <= 6; i++ {
		g.Go(func() error {
			n := ativas.Add(1)
			for {
				m := maximo.Load()
				if n <= m || maximo.CompareAndSwap(m, n) {
					break
				}
			}
			time.Sleep(100 * time.Millisecond)
			ativas.Add(-1)
			return nil
		})
	}
	g.Wait()
	fmt.Println("máximo simultâneo:", maximo.Load())
}
