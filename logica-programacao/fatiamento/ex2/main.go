package main

import "fmt"

func main() {
	letras := []string{"a", "b", "c", "d", "e", "f"}

	a := letras[2:5]
	b := letras[:3]
	c := letras[3:]
	d := letras[1:1]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
