package main

import "fmt"

func main() {
	fmt.Println("início")
	goto fim // linha 7: salta para "fim"
	x := 5   // linha 8: declaração PULADA
	fmt.Println("x vale", x)
fim:
	fmt.Println("fim") // aqui, x estaria "no escopo" sem ter sido criado
}
