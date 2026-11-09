package main

import "fmt"

//A regra para pertencer ao conjunto SIM é: a string deve conter pelo menos um caractere não-ASCII (isto é, pelo menos uma rune com valor Unicode maior que 127 / código além do ASCII padrão de 7 bits, ou que ocupe mais de 1 byte em UTF-8).

func ContemNaoASCII(s string) bool {
	for _, r := range s {
		if r >= 127 {
			return true
		}
	}
	return false

}

func main() {
	fmt.Println(ContemNaoASCII("Olá"))
	fmt.Println(ContemNaoASCII("ola"))
	fmt.Println(ContemNaoASCII("ç"))
	fmt.Println(ContemNaoASCII("cafe"))
	fmt.Println(ContemNaoASCII("世界"))
	fmt.Println(ContemNaoASCII(""))
	fmt.Println(ContemNaoASCII("café"))
	fmt.Println(ContemNaoASCII("x1"))
}
