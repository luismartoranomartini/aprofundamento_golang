package main

import (
	"fmt"
)

func mostra(nome string, s []int) {
	fmt.Printf("%-12s %-16v len=%d cap=%d\n", nome, fmt.Sprint(s), len(s), cap(s))
}

func main() {
	a := []int{10, 20, 30, 40, 50}

	mostra("a[1:3]", a[1:3])
	mostra("a[1:3:4]", a[1:3:4])
	mostra("a[1:3:#]", a[1:3:3])
	mostra("a[:0:0]", a[:0:0])
}
