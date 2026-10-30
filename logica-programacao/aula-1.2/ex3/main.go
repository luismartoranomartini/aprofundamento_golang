package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	maior := a

	if b > maior {
		maior = b
	}

	if c > maior {
		maior = c
	}

	fmt.Printf("O maior número entre %d, %d e %d é: %d\n", a, b, c, maior)

	n1, n2 := 7, 8
	if n1 > n2 {
		fmt.Printf("%d é maior que %d\n", n1, n2)
	} else if n1 == n2 {
		fmt.Printf("%d e %d são iguais\n", n1, n2)
	} else {
		fmt.Printf("%d é menor que %d\n", n1, n2)
	}
}
