package main

import "fmt"

func main() {
	minutos := 720 // Exemplo: 720 / 60 = 12

	hora := minutos / 60

	switch {
	case hora < 6:
		fmt.Println("madrugada")
	case hora < 12:
		fmt.Println("manhã")
	case hora < 18:
		fmt.Println("tarde")
	case hora < 24:
		fmt.Println("noite")
	default:
		fmt.Println("minutos inválidos")
	}
}
