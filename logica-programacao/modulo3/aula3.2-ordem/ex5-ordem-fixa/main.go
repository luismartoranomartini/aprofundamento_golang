package main

import "fmt"

func main() {
	notas := map[string]int{"Caio": 6, "Ana": 8, "Bia": 10}

	ordem := []string{"Ana", "Bia", "Caio"}

	for _, nome := range ordem {
		fmt.Println(nome, notas[nome])
	}
}
