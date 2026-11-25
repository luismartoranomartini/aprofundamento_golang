package main

import "fmt"

type Aluno struct {
	Nome string
	Nota float64
}

func main() {
	a := Aluno{"Ana", 8.5}
	fmt.Println(a.Idade)
}
