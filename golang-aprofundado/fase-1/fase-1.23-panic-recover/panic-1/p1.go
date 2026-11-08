package main

import "fmt"

func b() {
	defer fmt.Println("defer de b")
	panic("deu ruim")
	fmt.Println("nunca chega aqui")
}

func a() {
	defer fmt.Println("defer de a")
	b()
	fmt.Println("a: nunca chega aqui")
}

func main() {
	defer fmt.Println("defer de main")
	a()
	fmt.Println("main: nunca chega aqui")
}
