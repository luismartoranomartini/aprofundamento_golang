package main

import "fmt"

func main() {
	// make([]int, 3, 5): len 3, cap 5. Sobram 2 posições livres no array.
	s := make([]int, 3, 5)

	// Cabe (len 4 <= cap 5): o append escreve no MESMO array.
	t := append(s, 4)
	fmt.Println(len(t), cap(t), &s[0] == &t[0]) // 4 5 true

	// Cabe de novo (len 5 <= cap 5): ainda o mesmo array.
	u := append(t, 5)
	fmt.Println(len(u), cap(u), &s[0] == &u[0]) // 5 5 true

	// Não cabe (len 6 > cap 5): o Go aloca um array NOVO, copia tudo e o endereço muda.
	v := append(u, 6)
	fmt.Println(len(v), cap(v) > 5, &s[0] == &v[0]) // 6 true false
}
