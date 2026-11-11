package main

import "fmt"

func main() {
	s := "Golang"

	fmt.Println(s[3:3] == "")
	fmt.Println(s[len(s):] == "")
}
