package main

import "fmt"

func main() {
	fmt.Println(7/2, 7%2)     // 3 1
	fmt.Println(-7/2, -7%2)   // -3 -1
	fmt.Println(7/-2, 7%-2)   // -3 1
	fmt.Println(-7/-2, -7%-2) // 3 -1

	// func ehImparErrado(n int) bool { return n%2 == 1 }
	// func ehImpar(n int) bool       { return n%2 != 0 }

	// ehImparErrado(3)   // true
	// ehImparErrado(-3)  // false  ← ERRADO: -3 % 2 vale -1, e não 1
	// ehImpar(3)         // true
	// ehImpar(-3)        // true   ← certo

	n := -123
	soma := 0
	for n != 0 { // com "n > 0" o laço nem rodaria para negativos
		soma += n % 10 // -3, depois -2, depois -1
		n /= 10
	}
	fmt.Println(soma) // -6  ← os dígitos saem negativos

	m := -123
	if m < 0 {
		m = -m // agora m vale 123
	}
	s2 := 0
	for m > 0 {
		s2 += m % 10
		m /= 10
	}
	fmt.Println(s2) // 6

	fmt.Println(-9/2, -9%2)
	fmt.Println(9/-2, 9%-2)
}
