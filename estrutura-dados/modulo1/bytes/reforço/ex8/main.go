package main

import "fmt"

func main() {
	palavra := "banana"
	bytePalavra := []byte(palavra)

	for i := 0; i < len(bytePalavra); i++ {
		if bytePalavra[i] == 'a' {
			bytePalavra[i] += 'o' - 'a'
		}
	}
	palavra = string(bytePalavra)
	fmt.Println(palavra)
}
