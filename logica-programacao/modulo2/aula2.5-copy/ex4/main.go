package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	n := copy(s[1:], s)
	fmt.Println("para a direita:", n, s)

	t := []int{1, 2, 3, 4, 5}
	m := copy(t, t[1:])
	fmt.Println("para a esquerda:", m, t)
}
