package main

import "fmt"

func muda(a *int, b int) {
	*a = *a + b
	b = b + 100
	a = &b
	*a = *a + 1
	fmt.Println(*a, b)
}

func main() {
	x, y := 1, 10
	muda(&x, y)
	fmt.Println(x, y)
}
