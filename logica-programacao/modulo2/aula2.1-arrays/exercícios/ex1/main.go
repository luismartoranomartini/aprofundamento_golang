package main

import "fmt"

func somar(a [4]int) {
	soma := 0
	for _, s := range a {
		soma += s
	}
	fmt.Println(soma)
}

func main() {
	numeros := [4]int{10, 20, 30, 40}
	somar(numeros)
}
