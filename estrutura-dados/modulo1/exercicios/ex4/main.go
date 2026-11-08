package main

import (
	"fmt"
)

func main() {
	frase := "ir para casa agora"

	espacos := 0

	for i := 0; i < len(frase); i++ {
		if frase[i] == ' ' {
			espacos++
		}
	}
	fmt.Println(espacos)
}
