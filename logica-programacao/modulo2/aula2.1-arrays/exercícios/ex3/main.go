package main

import "fmt"

func inverte(a [5]int) [5]int {
	var resultado [5]int
	for i := 0; i < len(a); i++ {
		j := len(a) - 1 - i
		resultado[i] = a[j]
	}
	return resultado
}

func main() {
	numeros := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Invertido: ", inverte(numeros))
	fmt.Println("original: ", numeros)

}
