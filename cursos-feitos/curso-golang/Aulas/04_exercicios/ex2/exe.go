package main

import "fmt"

func main() {

	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	soma := 0
	soma1 := 0
	lista1 := lista[0:5]
	lista2 := lista[5:9]

	for i := 0; i < len(lista1); i++ {
		soma += lista1[i]
	}
	fmt.Println(soma)

	for i := 0; i < len(lista2); i++ {
		soma1 += lista2[i]
	}
	fmt.Println(soma1)

}
