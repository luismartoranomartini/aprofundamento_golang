package main

import "fmt"

func zera(s []int) {
	for i := range s {
		s[i] = 0
	}
	fmt.Println(s)
}

func main() {
	n := []int{1, 2, 3, 4}
	zera(n)
	fmt.Println(n)
}

/*Em Go, tudo é passado por valor (cópia). Porém, um slice não guarda os dados diretamente dentro
* de si — é uma estrutura pequena (chamada de slice header) composta por três campos:
1. Ponteiro;
2. len;
3. cap;
*/
