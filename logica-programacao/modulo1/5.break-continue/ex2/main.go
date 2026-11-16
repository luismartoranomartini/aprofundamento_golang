package main

import "fmt"

func main() {
	frase := "ir para casa Luís"
	letras := 0

	for i := 0; i < len(frase); i++ {
		if frase[i] == ' ' {
			continue
		}
		letras++
	}
	fmt.Println(letras)
}
