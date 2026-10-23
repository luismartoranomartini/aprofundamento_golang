package main

import "fmt"

func classifica(n int) string {
	if n < 0 {
		return "negativo"
	}
	if n == 0 {
		return "zero"
	}
	return "positivo"
}

func main() {
	fmt.Println(classifica(-1))
	fmt.Println(classifica(0))
	fmt.Println(classifica(9))
}
