package main

import "fmt"

func main() {
	var nota float64

	fmt.Print("Qual é a nota do aluno: ")
	fmt.Scan(&nota)

	switch {
	case nota >= 9:
		fmt.Println("A")
	case nota >= 7:
		fmt.Println("B")
	case nota >= 5:
		fmt.Println("C")
	default:
		fmt.Println("D")

	}
}
