package main

import "fmt"

func terminaComPonto(s string) bool {
	if s == "" {
		return false
	}
	return s[len(s)-1] == '.'
}

func main() {

	s := "Eu quero ver minha mãe."
	fmt.Println(terminaComPonto(s))

	s1 := ""
	fmt.Println(terminaComPonto(s1))
}
