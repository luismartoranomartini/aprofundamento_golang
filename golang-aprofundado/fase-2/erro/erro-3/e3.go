package main

import "fmt"

type ErroValidacao struct {
	Campo string
	Valor int
}

func (e *ErroValidacao) Error() string {
	return fmt.Sprintf("campo %q inválido: %d", e.Campo, e.Valor)
}

func valida(idade int) error {
	if idade < 0 {
		return &ErroValidacao{"idade", idade}
	}
	return nil
}

func main() {
	err := valida(-5)
	fmt.Println(err)

	if ve, ok := err.(*ErroValidacao); ok {
		fmt.Println("campo:", ve.Campo, "valor:", ve.Valor)
	}
	fmt.Println(valida(30) == nil)
}
