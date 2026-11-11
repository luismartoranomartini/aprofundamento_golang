package main

import (
	"fmt"
	"unicode"
)

func main() {
	letras, digitos := 0, 0

	for _, r := range "R2-D2 é ótimo" {
		if unicode.IsLetter(r) {
			letras++
		} else if unicode.IsDigit(r) {
			digitos++
		}
	}
	fmt.Println(letras, digitos)
}
