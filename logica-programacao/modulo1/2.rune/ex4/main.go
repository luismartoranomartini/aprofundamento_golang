package main

import "fmt"

func main() {
	r := []rune("olá")
	fmt.Println(r)
	fmt.Printf("%c\n", r[0])
	fmt.Printf("%T\n", r[2])
}
