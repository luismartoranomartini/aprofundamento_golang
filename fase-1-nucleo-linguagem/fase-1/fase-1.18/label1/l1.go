package main

import "fmt"

func main() {
	// sem rótulo
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				break
			}
			fmt.Println("sem rótulo", i, j)
		}
	}

fora:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break fora
			}
			fmt.Println("break fora", i, j)
		}
	}
	fmt.Println("---")

externo:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue externo
			}
			fmt.Println("continue externo", i, j)
		}
	}
}
