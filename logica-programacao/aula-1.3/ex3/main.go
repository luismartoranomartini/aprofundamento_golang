package main

import "fmt"

func sacar(saldo, valor float64) float64 {
	if valor <= 0 {
		return saldo
	}
	if valor > saldo {
		return saldo
	}
	return saldo - valor
}

func main() {
	fmt.Println(sacar(100, 30))
	fmt.Println(sacar(100, 100))
	fmt.Println(sacar(100, 0))
	fmt.Println(sacar(100, -5))
}
