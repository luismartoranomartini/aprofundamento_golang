package main

import "fmt"

func main() {
	numeros := []int{1, 2, 1, 3, 1, 2}

	contador := make(map[int]int)

	for _, num := range numeros {
		contador[num]++
	}
	fmt.Println("1:", contador[1])
	fmt.Println("2:", contador[2])
	fmt.Println("3:", contador[3])
	fmt.Println("4:", contador[4])
}

// []int{1, 2, 1, 3, 1, 2}
// 1, 2, 3 e 4
