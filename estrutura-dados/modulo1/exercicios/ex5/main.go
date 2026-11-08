package main

import "fmt"

func main() {
	digitos := "999999999999999999999999999999"
	var soma int

	for i := 0; i < len(digitos); i++ {
		soma += int(digitos[i] - '0')
	}
	fmt.Println(soma)
}
