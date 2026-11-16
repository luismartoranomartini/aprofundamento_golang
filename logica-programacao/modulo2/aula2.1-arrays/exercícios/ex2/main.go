package main

import "fmt"

func main() {
	numero := [5]int{10, 2, 3, 4, 5}
	maior := 0
	for i := range numero {
		if numero[i] > maior {
			maior = numero[i]
		}
	}
	fmt.Println("O maior número é:", maior)
}
