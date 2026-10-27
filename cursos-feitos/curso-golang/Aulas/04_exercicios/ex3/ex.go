package main

import "fmt"

func main() {

	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	num5 := 0
	num10 := 0

	for i := 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			num5 += lista[i]
		} else {
			num10 += lista[i]

		}
	}

	fmt.Println(num5)
	fmt.Println(num10)
}