package main

import "fmt"

func main() {
	palavras := []string{"casa", "arara", "xx", "aa1", "gato", "x", "", "abc"}

	for _, palavra := range palavras {
		letrasRepetidas := false

		for i := 0; i < len(palavra); i++ {
			for j := i + 1; j < len(palavra); j++ {
				if palavra[i] == palavra[j] {
					letrasRepetidas = true
					break
				}
			}
			if letrasRepetidas {
				break
			}
		}
		if letrasRepetidas {
			fmt.Printf("%-6q -> SIM\n", palavra)
		} else {
			fmt.Printf("%-6q -> NÃO\n", palavra)
		}
	}

}
