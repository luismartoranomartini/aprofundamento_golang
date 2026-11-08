package main

import "fmt"

type Falante interface {
	Falar() string
}

type Gato struct{}

func (Gato) Falar() string { return "miau" }

func main() {
	var x any = "texto"

	s, ok := x.(string)
	fmt.Println(s, ok)

	n, ok := x.(int)
	fmt.Println(n, ok)

	x = Gato{}
	if f, ok := x.(Falante); ok {
		fmt.Println(f.Falar())
	} else {
		fmt.Println("x não fala")
	}

	var y any = "texto"
	fmt.Println(y.(int))
}
