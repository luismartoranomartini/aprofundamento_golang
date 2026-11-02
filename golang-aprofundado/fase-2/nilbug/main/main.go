package main

import "fmt"

type ErroValidacao struct {
	Campo string
}

func (e *ErroValidacao) Error() string {
	return "campo inválido: " + e.Campo
}

func validarBug(nome string) error {
	var e *ErroValidacao
	if nome == "" {
		e = &ErroValidacao{Campo: "nome"}
		return e
	}
	return e
}

func validarOK(nome string) error {
	var e *ErroValidacao
	if nome == "" {
		e = &ErroValidacao{Campo: "nome"}
		return e
	}
	return nil
}

func main() {
	e1 := validarBug("Ana")
	fmt.Println(e1 == nil)
	fmt.Printf("%T\n", e1)
	if e1 != nil {
		fmt.Println("entrou no if com nome válido")
	}

	e2 := validarOK("Ana")
	fmt.Println(e2 == nil)
	fmt.Printf("%T\n", e2)

	e3 := validarOK("")
	fmt.Println(e3 == nil, e3)
}
