package main

import "fmt"

func Pares(s []string) func(yield func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	for i, v := range Pares([]string{"a", "b"}) {
		fmt.Println(i, v)
	}
}
