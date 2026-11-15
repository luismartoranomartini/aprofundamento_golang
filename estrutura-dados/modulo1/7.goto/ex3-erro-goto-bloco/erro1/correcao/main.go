package main

import "fmt"

func main() {
	fmt.Println("início")
	x := 5   // linha 8: declaração antes de goto
	goto fim // linha 7: salta para "fim"
fim:
	fmt.Println("fim, x=", x) // aqui, x estaria "no escopo" sem ter sido criado
}
