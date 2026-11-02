package main

import "fmt"

type Caixa struct {
	Itens []int
}

func main() {
	a := Caixa{[]int{1}}
	b := Caixa{[]int{1}}
	fmt.Println(a == b)
}
