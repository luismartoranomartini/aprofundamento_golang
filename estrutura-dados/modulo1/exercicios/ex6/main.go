package main

import (
	"fmt"
)

func main() {
	palavra := "go lang"
	bytePalavra := []byte(palavra)

	for i := 0; i < len(palavra); i++ {
		bytePalavra[i] = palavra[i] - 'a' + 'A'
	}
	palavra = string(bytePalavra)
	fmt.Println(palavra)

	fmt.Println(32 - 'a' + 'z')

}
