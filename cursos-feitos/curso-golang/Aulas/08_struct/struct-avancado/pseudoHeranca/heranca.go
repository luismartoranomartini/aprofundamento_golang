package main

import "fmt"

type carro struct {
	nome string
	vel  int
}

type ferrari struct {
	carro
	turbo bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.vel = 0
	f.turbo = true
	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turbo)
	fmt.Println(f)
}