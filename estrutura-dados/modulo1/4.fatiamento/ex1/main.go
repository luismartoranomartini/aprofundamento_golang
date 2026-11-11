package main

import "fmt"

func main() {
	s := "Golang"

	fmt.Println(s[0:2])
	fmt.Println(s[2:])
	fmt.Println(s[:3])
	fmt.Println(s[:])
}
