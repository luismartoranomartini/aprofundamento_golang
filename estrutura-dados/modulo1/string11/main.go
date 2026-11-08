package main

import "fmt"

func main() {
	s := "Go"
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	fmt.Printf("%d\n", []byte(s))
}
