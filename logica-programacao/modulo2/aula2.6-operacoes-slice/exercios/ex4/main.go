package main

import "fmt"

// inverte inverte s no lugar, sem criar outro slice e sem devolver nada.
func inverte(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4}
	inverte(a)
	fmt.Println(a) // esperado: [4 3 2 1]

	b := []int{1, 2, 3}
	inverte(b)
	fmt.Println(b) // esperado: [3 2 1]

	var c []int
	inverte(c)
	fmt.Println(c) // esperado: []

	d := []int{7}
	inverte(d)
	fmt.Println(d) // esperado: [7]
}
