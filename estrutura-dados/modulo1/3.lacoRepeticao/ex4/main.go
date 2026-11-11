package main

import "fmt"

func main() {
	n := 12345
	digito := 0
	for n > 0 {
		digito++
		n /= 10
	}
	fmt.Println(digito)
}
