package main

import "fmt"

type Item struct {
	N int
}

func main() {
	s := []string{"a", "b", "c"}

	for i := range s {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()

	for i, v := range s {
		fmt.Print(i, v, " ")
	}
	fmt.Println()

	// v é uma cópia do elemento
	itens := []Item{{1}, {2}}
	for i := range itens {
		itens[i].N = 99
	}
	fmt.Println(itens)

	// nil slice: zero voltas
	var nada []int
	for range nada {
		fmt.Println("nunca")
	}
	fmt.Println("ok")
}
