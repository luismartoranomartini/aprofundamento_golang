package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operacaoLenta(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	ini := time.Now()
	err := operacaoLenta(ctx)
	fmt.Printf("erro: %v (depos de %.1fs)\n", err, time.Since(ini).Seconds())
	fmt.Println("é deadline?", errors.Is(err, context.DeadlineExceeded))
}
