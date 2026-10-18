package main

import "fmt"

type Ponto struct{ X, Y int }

type Aluno struct {
	Nome string
	Nota int
}

func main() {
	fmt.Println("1) struct como chave")
	fmt.Println()
	grade := map[Ponto]string{{1, 2}: "A", {3, 4}: "B"}
	fmt.Println(grade[Ponto{1, 2}], grade[Ponto{9, 9}] == "")

	fmt.Println("2) struct como valor: copiar, alterar, devolver")
	alunos := map[string]Aluno{"ana": {"Ana", 7}}
	a := alunos["ana"] // cópia
	a.Nota = 10        // altera a cópia
	alunos["ana"] = a  // devolve ao map
	fmt.Println(alunos)

	fmt.Println("3) map de ponteiros: altera direto")
	ptrs := map[string]*Aluno{"bia": {"Bia", 5}}
	ptrs["bia"].Nota = 9
	fmt.Println(*ptrs["bia"])

	fmt.Println("4) map de slices: append + reatribuir")
	tags := map[string][]string{}
	tags["go"] = append(tags["go"], "slice")
	tags["go"] = append(tags["go"], "map")
	fmt.Println(tags, len(tags["java"]))

	fmt.Println("5) map dentro de map: iniciar o interno")
	aninhado := map[string]map[string]int{}
	if aninhado["x"] == nil {
		aninhado["x"] = map[string]int{}
	}
	aninhado["x"]["y"] = 1
	fmt.Println(aninhado)

	m := map[string]map[string]int{}
	fmt.Println(m["x"]["y"])
}
