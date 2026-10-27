package main

import "fmt"

func imprimeSeguro(p *int) {
	if p == nil {
		fmt.Println("ponteiro nulo, nada a fazer")
		return
	}
	fmt.Println(*p)
}

func main() {
	var p *int
	fmt.Println(p) // nil
	fmt.Println(p == nil)
	// fmt.Println(*p) // nil
	imprimeSeguro(p)
}
