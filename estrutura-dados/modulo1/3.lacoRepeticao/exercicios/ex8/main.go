package main

import "fmt"

func main() {
	s := "abc"
	for i := 0; i <= len(s); i++ { // <= em vez de <
		fmt.Println(s[i])
	}
	/* 97
	98
	99
	panic: runtime error: index out of range [3] with length 3

	goroutine 1 [running]:
	main.main()
	        /run/media/luismmartini/Luis/Go/estrutura-dados/modulo1/3.lacoRepeticao/exercicios/ex8/main.go:8 +0x7a
	exit status 2 */
	//	CORREÇÃO
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
