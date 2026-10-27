package main

import "fmt"

func main() {
	var exemplo any
	fmt.Println(exemplo)
	fmt.Println()

	// ----
	var (
		numero  int = 5
		numero2 any = 5
	)

	numero = numero + 1
	numero++

	numero2Int, ok := numero2.(int)
	if ok {
		numero2Int++
	}

	// var (
	// 	lista  any = []int{1, 2, 3, 4, 5, 6}
	// 	lista2 any = []int{1, 2, 3, 4, 5, 6}
	// )
	// if lista == lista2 {
	//
	// }
	// ----
	var x any
	x = "texto"

	switch v := x.(type) {
	case int:
		fmt.Println("int\n", v)
	case string:
		fmt.Println("texto\n", v)
	}

}
