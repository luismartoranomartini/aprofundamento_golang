package main

import "fmt"

func main() {
	n := 9876543
	digitos := 0

	for n > 0 {
		digitos++
		n /= 10
	}
	fmt.Println(digitos)
}
