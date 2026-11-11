package main

import "fmt"

func main() {
	s := "4096"
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
