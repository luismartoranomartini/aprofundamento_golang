package main

import "fmt"

func main() {
	var mes int

	fmt.Print("Digite o número do mês (1 a 12): ")
	_, err := fmt.Scan(&mes)
	if err != nil {
		fmt.Println("Entrada inválida.")
		return
	}

	switch mes {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 dias")
	case 4, 6, 9, 11:
		fmt.Println("30 dias")
	case 2:
		fmt.Println("28 ou 29 dias")
	default:
		fmt.Println("Mes invalido")
	}
}
