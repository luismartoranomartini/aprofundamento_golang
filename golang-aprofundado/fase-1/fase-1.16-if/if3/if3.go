package main

import (
	"errors"
	"fmt"
)

func falha() (int, error) { return 0, errors.New("deu ruim") }

func main() {
	var err error
	if true {
		v, err := falha() // := cria err NOVO neste bloco
		fmt.Println("dentro:", v, err)
	}
	fmt.Println("fora:", err)

	var err2 error
	if true {
		var v int
		v, err2 = falha() // = usa o err2 de fora
		fmt.Println("dentro:", v, err2)
	}
	fmt.Println("fora:", err2)
}
