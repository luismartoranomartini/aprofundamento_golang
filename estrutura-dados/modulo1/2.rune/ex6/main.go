package main

import "fmt"

func main() {
	r := 'á'
	fmt.Printf("%d\n", r) // 225
	fmt.Printf("%c\n", r) // á
	fmt.Printf("%q\n", r) // 'á'
	fmt.Printf("%U\n", r) // U+00E1
	fmt.Printf("%x\n", r) // e1
	fmt.Printf("%T\n", r) // int32
}
