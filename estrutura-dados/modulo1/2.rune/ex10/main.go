package main

import "fmt"

func main() {
	s := "abc"
	r := 'a'
	fmt.Println(s[0] == r)
	fmt.Println(rune(s[0]) == r)
	fmt.Println(s[0] == 'a')
}
