package main

import "fmt"

func conceito(nota float64) string {
	if nota < 0 || nota > 10 {
		return "nota inválida"
	}
	if nota >= 7 {
		return "aprovado"
	} else if nota >= 5 {
		return "recuperação"
	} else {
		return "reprovado"
	}
}

func main() {
	fmt.Println(conceito(-1))
	fmt.Println(conceito(11))
	fmt.Println(conceito(4))
	fmt.Println(conceito(5))
	fmt.Println(conceito(7))
	fmt.Println(conceito(10))
}
