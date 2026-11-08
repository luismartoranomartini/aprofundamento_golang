package main

import (
	"errors"
	"fmt"
)

var ErrNaoAchou = errors.New("não achou")

func repositorio() error { return ErrNaoAchou }

func servico() error {
	if err := repositorio(); err != nil {
		return fmt.Errorf("serviço: %w", err)
	}
	return nil
}

func handler() error {
	if err := servico(); err != nil {
		return fmt.Errorf("handler: %w", err)
	}
	return nil
}

func main() {
	err := handler()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrNaoAchou))
	fmt.Println(errors.Is(err, errors.New("não achou")))
	for e := err; e != nil; e = errors.Unwrap(e) {
		fmt.Printf("%T → %v\n", e, e)
	}
}
