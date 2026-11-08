package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Contem[T comparable](s []T, alvo T) bool {
	for _, v := range s {
		if v == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Max(3, 7), Max("abc", "abd"), Max(2.5, 1.5))
	fmt.Println(Max[float64](1, 2.5))
	fmt.Println(Contem([]string{"a", "b"}, "b"), Contem([]int{1, 2}, 9))
	fmt.Println(max(3, 7))
	fmt.Println(max("abc", "abd"))
	fmt.Println(max(2.5, 1.5))
	fmt.Println(max(1, 2.5))
}
