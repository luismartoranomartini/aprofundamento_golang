package main

import "fmt"

func gerador(passo int) func() int {
	n := 0
	return func() int {
		n += passo
		return n
	}
}

func main() {
	a := gerador(2)
	b := gerador(5)
	fmt.Println(a(), a())
	fmt.Println(b(), b())
	fmt.Println(a(), b())
}
