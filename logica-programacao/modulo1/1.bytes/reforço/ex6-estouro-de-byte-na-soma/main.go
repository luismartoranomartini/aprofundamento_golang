package main

import "fmt"

func main() {
	s := "9999999999999999999999999999999999999999"

	var somaByte byte = 0
	var somaInt int = 0

	for i := 0; i < len(s); i++ {
		num := s[i]
		if num >= '0' && num <= '9' {
			valorDigito := num - '0'
			somaByte += valorDigito
			somaInt += int(valorDigito)
		}
	}

	fmt.Println("Byte:", somaByte)
	fmt.Println("Int:", somaInt)
}
