package main

import "fmt"

func palavra(s string) {

	inicio := -1

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && inicio == -1 {
			inicio = i
		}
		if s[i] == ' ' && inicio != -1 {
			fmt.Printf("%q ", s[inicio:i])
			inicio = -1
		}
	}
	if inicio != -1 {
		fmt.Printf("%q ", s[inicio:])
	}
}

func main() {
	frase1 := "Luís Martini"
	frase2 := "Luís"
	frase3 := "Luís Francisco Martorano Martini"

	palavra(frase1)
	palavra(frase2)
	palavra(frase3)
}
