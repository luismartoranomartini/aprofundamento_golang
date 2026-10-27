package main

import "fmt"

func main() {
	var temp float64

	fmt.Print("Digite uma temperatura: ")
	fmt.Scan(&temp)

	switch {
	case temp < 0:
		fmt.Println("Congelante")
	case temp < 15:
		fmt.Println("Frio")
	case temp < 25:
		fmt.Println("Agradável")
	case temp < 35:
		fmt.Println("Quente")
	default:
		fmt.Println("Muito quente")
	}
}
