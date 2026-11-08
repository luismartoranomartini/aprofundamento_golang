package main

import "fmt"

func main() {
	s := "banaba"
	fmt.Printf("%c\n", s[len(s)-1])

	v := ""
	fmt.Println(len(v), v == "")
}
