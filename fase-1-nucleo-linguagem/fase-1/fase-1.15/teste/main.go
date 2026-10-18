package main

import "fmt"

func Dobros(n int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			fmt.Println("produz", i*2)
			if !yield(i * 2) {
				return
			}
		}
	}
}

func main() {
	for v := range Dobros(4) {
		fmt.Println("usa", v)
		if v == 4 {
			break
		}
	}
	fmt.Println("fim")
}
