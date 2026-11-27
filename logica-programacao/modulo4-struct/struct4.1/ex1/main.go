package main

import "fmt"

type Aluno struct {
	Nome string
	Nota float64
}

func main() {
	a := Aluno{Nome: "Ana", Nota: 8.5}

	b := Aluno{"Leo", 7}

	var c Aluno

	fmt.Println(a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
}
