package main

import (
	"fmt"
	"iter"
)

func Contar(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	var s iter.Seq[int] = Contar(3)
	for v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
