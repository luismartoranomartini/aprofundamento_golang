package main

import "fmt"

func main() {
	var s []int
	t := []int{}

	fmt.Println(len(s), cap(s))
	fmt.Println(len(t), cap(t))
	fmt.Println(s == nil)
	fmt.Println(t == nil)
}
