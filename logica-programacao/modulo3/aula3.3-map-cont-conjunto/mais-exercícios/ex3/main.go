package main

import "fmt"

func primeiroRepetido(s []int) (int, bool) {

	repetido := make(map[int]int)
	duasVezes := 0

	for _, n := range s {
		repetido[n]++
		if repetido[n] > 1 {
			duasVezes = n
			return duasVezes, true
		}
	}

	return 0, false
}

func main() {
	fmt.Println(primeiroRepetido([]int{3, 1, 4, 1, 5, 3}))
	fmt.Println(primeiroRepetido([]int{1, 2, 3}))
	fmt.Println(primeiroRepetido([]int{}))
}
