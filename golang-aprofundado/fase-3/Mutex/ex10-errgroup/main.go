package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func tarefa(ctx context.Context, nome string, dur time.Duration, falha bool) error {
	select {
	case <-time.After(dur):
		if falha {
			fmt.Println(nome, "falhou")
			return errors.New(nome + ": deu certo")
		}
		fmt.Println(nome, "terminou")
		return nil
	case <-ctx.Done():
		fmt.Println(nome, "cancelada:", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error { return tarefa(ctx, "A", 100*time.Millisecond, true) })
	g.Go(func() error { return tarefa(ctx, "B", 2*time.Second, false) })
	g.Go(func() error { return tarefa(ctx, "C", 2*time.Second, false) })

	err := g.Wait()
	fmt.Println("Wait devolveu:", err)
}
