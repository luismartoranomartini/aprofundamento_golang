package main

import "fmt"

func main() {
	s := "á"
	fmt.Println(len(s))
	fmt.Println([]byte(s))
	fmt.Println([]rune(s))
}
