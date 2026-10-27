package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	_ constraints.Float
	_ constraints.Integer
	_ constraints.Signed
	_ constraints.Unsigned
	_ constraints.Complex
	_ constraints.Ordered
	_ cmp.Ordered
)

func Min[T cmp.Ordered](v1, v2 T) T {
	var min T

	if v1 < v2 {
		return v1
	} else if v2 < v2 {
		return v2
	}
	return min
}

func main() {
	fmt.Println(Min(10, 100))
	fmt.Println(Min("a", "b"))
}
