package main

import "fmt"

func insere(s []int, i, v int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

func main() {
	s := []int{10, 20, 30, 40}

	s = insere(s, 0, 5)
	fmt.Println("inicio:", s)

	s = insere(s, 3, 25)
	fmt.Println("meio:", s)

	s = insere(s, len(s), 99)
	fmt.Println("fim:", s)

	var vazio []int
	vazio = insere(vazio, 0, 7)
	fmt.Println("vazio:", vazio)
}
