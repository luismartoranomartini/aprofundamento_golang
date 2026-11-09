package main

import (
	"fmt"
	"unicode"
)

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func main() {
	nome1 := "éric"
	nome2 := "ana"
	obj1 := "1a"
	vazio := ""

	fmt.Println(capitalize(nome1))
	fmt.Println(capitalize(nome2))
	fmt.Println(capitalize(obj1))
	fmt.Println(capitalize(vazio))
}
