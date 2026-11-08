package main

import "fmt"

func main() {
	var a, b any = 1, 1
	fmt.Println(a == b)
	var c, d any = 1, "1"
	fmt.Println(c == d)

	var s1, s2 any = []int{1}, []int{1}
	fmt.Println(s1 == s2)
}
