package main

import "fmt"

func Contar(n int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			fmt.Println(" [produzido", i, "]")
			if !yield(i) {
				fmt.Println(" [parou]")
				return
			}
		}
	}
}

func main() {
	for v := range Contar(3) {
		fmt.Println("recebi", v)
	}
	fmt.Println("---- con break")
	for v := range Contar(5) {
		fmt.Println("recebi", v)
		if v == 2 {
			break
		}
	}
	fmt.Println("fim")
}
