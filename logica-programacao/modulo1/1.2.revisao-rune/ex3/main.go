package main

import "fmt"

func main() {
	r := []rune("ação")
	b := []byte("ação")

	fmt.Printf("%c %d\n", r[2], r[2])
	fmt.Printf("%c %d\n", b[2], b[2])
}
