package main

import "fmt"

func main() {
externo:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				continue externo
			}
			if i == 1 {
				break externo
			}
			fmt.Println(i, j)
		}
		fmt.Println("fim do i", i)
	}
	fmt.Println("fim")
}
