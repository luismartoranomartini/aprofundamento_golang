package main

import "fmt"

func main() {
	s := "ir para casa agora"
	palavras := 0
	dentro := false
	for i := 1; i < len(s); i++ {
		if s[i] != ' ' && !dentro {
			palavras++
		}
		dentro = s[i] != ' '
	}
	fmt.Println(palavras)
}
