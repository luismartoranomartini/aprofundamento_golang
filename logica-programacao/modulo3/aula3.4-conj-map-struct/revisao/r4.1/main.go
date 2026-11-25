package main

import "fmt"

func soSeRepetem(numeros []int) []int {
	contador := make(map[int]int)
	var guarda []int

	for _, num := range numeros {
		contador[num]++
	}

	for _, num := range numeros {
		if contador[num] > 1 {
			guarda = append(guarda, num)
			contador[num] = 0
		}
	}
	return guarda
}

func main() {
	entrada := []int{1, 2, 2, 3, 3, 3, 4}
	fmt.Println(soSeRepetem(entrada))
}
