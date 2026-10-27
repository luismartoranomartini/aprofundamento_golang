package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

// type Number (
// 	_ constraints.Float
// 	_ constraints.Integer
// 	_ constraints.Signed
// 	_ constraints.Unsigned
// 	_ constraints.Ordered
// 	_ cmp.Ordered
// )

type Number interface {
	constraints.Integer | constraints.Float
}

// type Number constraints.Integer

func Soma[T Number](numeros []T) T {
	var soma T = 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func Min[T cmp.Ordered](v1, v2 T) T {
	var min T

	if v1 < v2 {
		return v1
	} else if v2 < v1 {
		return v2
	}
	return min
}

func main() {
	numeroInteiro := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numeroInteiro1 := []int32{1, 2, 3}
	numeroInteiro2 := []int{1, -2, 3}

	fmt.Println(Soma(numeroInteiro))
	fmt.Println(Soma(numeroInteiro1))
	fmt.Println(Soma(numeroInteiro2))

	fmt.Println(Min(500, 511))
	fmt.Println(Min("Luís", "Xoxão"))
}
