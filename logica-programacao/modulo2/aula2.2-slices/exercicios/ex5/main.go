package main

import "fmt"

func main() {
	s := make([]int, 2, 5)
	fmt.Println(s[3])

}

/*
Em Go, os limites de indexação direta (s[i]) são verificados exclusivamente contra o len (tamanho), nunca contra o cap (capacidade)
*/
