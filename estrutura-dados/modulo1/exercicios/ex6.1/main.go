package main

import "fmt"

func main() {
	palavra := "go lang"
	bytePalavra := []byte(palavra)

	for i := 0; i < len(palavra); i++ {
		c := palavra[i]
		if c >= 'a' && c <= 'z' {
			bytePalavra[i] = c - 'a' + 'A'
		}
	}
	palavra = string(bytePalavra)
	fmt.Printf("%q\n", palavra)
}
