package main

import "fmt"

func main() {
	// 1) clássico
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2) só condição
	n := 1
	for n < 50 {
		n *= 3
	}
	fmt.Println(n)

	// 3) infinito, com break
	c := 0
	for {
		c++
		if c == 4 {
			break
		}
	}
	fmt.Println(c)

	// 4) range sobre inteiro
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5) range sem usar a variável
	k := 0
	for range 3 {
		k++
	}
	fmt.Println(k)

	// 6) range com n <= 0:: zero iterações
	for i := range -2 {
		fmt.Println("nunca", i)
	}
	fmt.Println("fim")
}
