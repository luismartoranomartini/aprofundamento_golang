package main

import "fmt"

func main() {
	s := "ação"
	fmt.Println(len(s), len([]rune(s)))
}
