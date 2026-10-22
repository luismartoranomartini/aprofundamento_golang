package main

import "fmt"

func troca(p *int) {
	outra := 999
	p = &outra
	*p = 1
}

func escreve(p *int) {
	*p = 1
}

func main() {
	x := 5
	px := &x

	troca(px)
	fmt.Println("depois de troca: x =", x, "| px ainda aponta para x:", px == &x)

	escreve(px)
	fmt.Println("depois de escreve: x =", x)
}
