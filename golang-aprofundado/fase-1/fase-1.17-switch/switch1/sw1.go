package main

import "fmt"

func dia(n int) string {
	switch n {
	case 0, 6:
		return "fim de semana"
	case 1, 2, 3, 4, 5:
		return "dia útil"
	default:
		return "inválido"
	}
}

func main() {
	fmt.Println(dia(0), "|", dia(3), "|", dia(9))

	switch 2 {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch x := 7 * 2; x {
	case 14:
		fmt.Println("quatroze", x)
	}

	nota := 6
	switch {
	case nota >= 7:
		fmt.Println("aprovado")
	case nota >= 5:
		fmt.Println("recuperação")
	default:
		fmt.Println("reprovado")

	}
	switch 99 {
	case 1:
		fmt.Println("um")
	}
	fmt.Println("fim")
}
