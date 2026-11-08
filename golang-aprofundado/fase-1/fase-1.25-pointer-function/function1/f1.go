package main

import "fmt"

func dobraValor(n int) {
	n = n * 2
}

func dobraPonteiro(n *int) {
	*n = *n * 2
}

func main() {
	x := 5
	dobraValor(x)
	fmt.Println("depois de dobraValor: ", x)
	dobraPonteiro(&x)
	fmt.Println("depois de dobraPonteiro: ", x)
}
