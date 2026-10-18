package main

import "fmt"

func main() {
	s := "olá"

	b := []byte(s)
	r := []rune(s)
	fmt.Println("byte: ", b, len(b))
	fmt.Println("rune: ", r, len(r))

	b[0] = 'O'
	fmt.Println(string(b), s)

	r[2] = 'e'
	fmt.Println(string(r), s)

	fmt.Println(string(b) == "Olá")
}
