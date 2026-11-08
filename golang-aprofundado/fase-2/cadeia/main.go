package main

import (
	"errors"
	"fmt"
)

var ErrNaoEncontrado = errors.New("não encontrado")

type ErroConsulta struct {
	ID int
}

func (e ErroConsulta) Error() string {
	return fmt.Sprintf("consulta do id %d falhou", e.ID)
}

func repositorio(id int) error {
	if id == 0 {
		return ErrNaoEncontrado
	}

	if id < 0 {
		return ErroConsulta{
			ID: id,
		}
	}
	return nil
}

func servico(id int) error {
	err := repositorio(id)
	if err != nil {
		return fmt.Errorf("servico: %w", err)
	}
	return nil
}

func handler(id int) error {
	err := servico(id)
	if err != nil {
		return fmt.Errorf("handler: %w", err)
	}
	return nil
}

func main() {
	fmt.Println(handler(5) == nil)

	e1 := handler(0)
	fmt.Println(e1)
	fmt.Println(errors.Is(e1, ErrNaoEncontrado))

	e2 := handler(-7)
	fmt.Println(e2)
	var ec ErroConsulta
	fmt.Println(errors.As(e2, &ec), ec.ID)
	fmt.Println(errors.Is(e2, ErrNaoEncontrado))
	fmt.Println(errors.Is(e2, ErroConsulta{ID: -7}))
	fmt.Println(errors.Unwrap(e2))
}
