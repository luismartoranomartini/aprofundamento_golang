package main

import "fmt"

func main() {
	s := "olá"
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %d (%#x)\n", i, s[i], s[i])
	}
	fmt.Printf("%c", s[2])
}
