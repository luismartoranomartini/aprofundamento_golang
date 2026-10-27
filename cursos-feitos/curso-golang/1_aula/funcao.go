package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(2, 3)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	f("Texto 1")
	txt := f("texto 2")
	fmt.Println(txt)

	resultadoSoma, resultadoSub := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSub)
}
