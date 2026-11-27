package main

import "fmt"

type Turma struct {
	Nome  string
	Notas []int
}

func main() {
	t1 := Turma{"A", []int{7, 8, 9}}

	t2 := t1
	t2.Nome = "B"
	t2.Notas[0] = 99
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)

	t3 := Turma{Nome: t1.Nome, Notas: make([]int, len(t1.Notas))}
	copy(t3.Notas, t1.Notas)
	t3.Notas[0] = 0
	fmt.Println("t1:", t1)
	fmt.Println("t3:", t3)

	turmas := []Turma{{"X", []int{1, 2}}, {"Y", []int{3, 4}}}
	for _, t := range turmas {
		t.Nome = "?"
		t.Notas[0] = 0
	}
	fmt.Println(turmas)
}
