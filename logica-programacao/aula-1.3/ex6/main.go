package main

import "fmt"

// Um número de três dígitos é especial se o dígito do meio for igual à soma do primeiro dígito com o último dígito.
func especial(n int) bool {
	if n < 100 || n > 999 {
		return false
	}

	primeiro := n / 100
	meio := (n / 10) % 10
	ultimo := n % 10
	return meio == primeiro+ultimo
}

func main() {
	fmt.Println(especial(121))
	fmt.Println(especial(242))
	fmt.Println(especial(330))
	fmt.Println(especial(462))
	fmt.Println(especial(123))
	fmt.Println(especial(135))
	fmt.Println(especial(100))
	fmt.Println(especial(555))
}
