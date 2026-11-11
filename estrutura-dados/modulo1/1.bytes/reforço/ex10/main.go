package main

import "fmt"

func main() {
	palavra := "luis"
	bytePalavra := []byte(palavra)

	if palavra == "" {
		fmt.Println("Não existe palavra")
		return
	}

	if bytePalavra[0] >= 'a' && bytePalavra[0] <= 'z' {
		bytePalavra[0] -= 'a' - 'A'
	}
	palavra = string(bytePalavra)
	fmt.Println(palavra)
}
