package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("falhou")
	e2 := errors.New("falhou")
	fmt.Println(e1 == e2, e1 == e1)
	fmt.Println(e1.Error(), e1)

	e3 := fmt.Errorf("código %d inválido", 42)
	fmt.Println(e3)
	fmt.Printf("%T %T\n", e1, e3)

	var nenhum error
	fmt.Println(nenhum == nil, nenhum)
}
