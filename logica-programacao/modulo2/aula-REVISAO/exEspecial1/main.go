package main

import "fmt"

func insere(s []int, i, v int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] += v
	return s
}

func main() {
	s := []int{10, 20, 30, 40}

	s = insere(s, 0, 5)       // no início
	fmt.Println("início:", s) // [5 10 20 30 40]

	s = insere(s, 3, 25)    // no meio
	fmt.Println("meio:", s) // [5 10 20 25 30 40]

	s = insere(s, len(s), 99) // no fim
	fmt.Println("fim:", s)    // [5 10 20 25 30 40 99]

	// Se o resultado não fosse reatribuído a s, o len do s da main não mudaria.
	t := []int{1, 2, 3}
	insere(t, 0, 7)
	fmt.Println(t) // [1 2 3] (resultado jogado fora)
}
