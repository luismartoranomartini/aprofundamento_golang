package main

import "fmt"

func main() {
	n, soma := 4096, 0
	for n > 0 {
		soma += n % 10
		n /= 10
	}
	fmt.Println(soma)
}
