package main

import (
	"errors"
	"fmt"
)

type ErroValidacao struct{ Campo string }

func (e *ErroValidacao) Error() string { return "inválido: " + e.Campo }

func valida() error { return &ErroValidacao{"idade"} }

func main() {
	err := fmt.Errorf("cadastro: %w", valida())

	_, ok := err.(*ErroValidacao)
	fmt.Println("assertion direta:", ok)

	var ve *ErroValidacao
	if errors.As(err, &ve) {
		fmt.Println("As:", ve.Campo)
	}
	fmt.Println(err)
}
