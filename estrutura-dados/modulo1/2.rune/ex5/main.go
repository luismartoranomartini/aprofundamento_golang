package main

import "fmt"

func main() {
	r := []rune("pão")
	r[0] = 'm'
	fmt.Println(string(r))
}
