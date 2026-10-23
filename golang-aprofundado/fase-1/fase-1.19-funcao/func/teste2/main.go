package main

import "fmt"

func muda(n int, s []int) {
	n = 50
	s[1] = 50
}

func main() {
	n := 5
	s := []int{1, 2, 3}
	muda(n, s)
	fmt.Println(n)
	fmt.Println(s)
}
