package main

import "fmt"

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func removerSemOrdem(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	orig := []int{1, 2, 3, 4, 5}
	t := remove(orig, 1)
	fmt.Println("remove(orig, 1):", t, "len=", len(t), "cap =", cap(t))
	fmt.Println("orig depois:", orig)

	u := []int{1, 2, 3, 4, 5}
	u = removerSemOrdem(u, 1)
	fmt.Println("sem ordem:", u)

	w := []int{1, 2, 3, 4, 5}
	w = append(w[:1], w[2:]...)
	fmt.Println("forma curta:", w)
}
