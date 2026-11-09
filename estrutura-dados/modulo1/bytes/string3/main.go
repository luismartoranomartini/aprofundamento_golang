package main

import "fmt"

func main() {
	s := "banana"
	fmt.Printf("%c\n", s[len(s)-1])

	v := ""
	fmt.Println(len(v), v == "")
}
