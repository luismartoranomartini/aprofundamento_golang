package main

import (
	"errors"
	"fmt"
)

func main() {
	var ErrNaoAchou = errors.New("não achou")

	comW := fmt.Errorf("buscar usuário: %w", ErrNaoAchou)
	semW := fmt.Errorf("buscar usuário: %v", ErrNaoAchou)

	fmt.Println(comW)
	fmt.Println(semW)
	fmt.Println("== direto:", comW == ErrNaoAchou)
	fmt.Println("Is com %w:", errors.Is(comW, ErrNaoAchou))
	fmt.Println("Is com %v:", errors.Is(semW, ErrNaoAchou))
	fmt.Println("Unwrap:", errors.Unwrap(comW), errors.Unwrap(comW) == ErrNaoAchou)
	fmt.Println("Unwrap sem w:", errors.Unwrap(semW))
}
