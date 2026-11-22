package main

import "fmt"

func soma(m map[string]int) int {
	valor := 0

	for _, v := range m {
		valor += v
	}

	return valor
}

func main() {
	somatorio := map[string]int{"a": 3, "b": 4}
	somatorio1 := make(map[string]int)
	var somatorio2 map[string]int

	resultado := soma(somatorio)
	resultado1 := soma(somatorio1)
	resultado2 := soma(somatorio2)

	fmt.Println(resultado)
	fmt.Println(resultado1)
	fmt.Println(resultado2)
}
