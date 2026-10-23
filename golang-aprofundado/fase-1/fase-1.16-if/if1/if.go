package main

import (
	"fmt"
	"strconv"
)

func main() {
	nota := 7
	if nota >= 7 {
		fmt.Println("aprovado")
	} else if nota >= 5 {
		fmt.Println("recuperação")
	} else {
		fmt.Println("reprovado")
	}

	// declaração inicial; condição
	if n, err := strconv.Atoi("42"); err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("número:", n) // n e err existem também no else
	}

	if _, err := strconv.Atoi("4x2"); err != nil {
		fmt.Println("erro:", err)
	}

	// comma-ok com map
	preco := map[string]int{"cafe": 8}
	if p, ok := preco["cafe"]; ok {
		fmt.Println("café custa", p)
	}
	if _, ok := preco["cha"]; !ok {
		fmt.Println("sem chá")
	}
}
