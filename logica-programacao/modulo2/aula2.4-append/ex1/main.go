package main

import "fmt"

func mostra(nome string, s []int) {
	fmt.Printf("%-8s %v len=%d cap=%d nil=%v\n", nome, s, len(s), cap(s), s == nil)
}

func main() {
	var s []int
	mostra("início", s)

	s = append(s, 18)
	mostra("+10", s)

	s = append(s, 20)
	mostra("+20", s)

	s = append(s, 30)
	mostra("+30", s)

	s = append(s, 40)
	mostra("+40", s)

	s = append(s, 50)
	mostra("+50", s)
}
