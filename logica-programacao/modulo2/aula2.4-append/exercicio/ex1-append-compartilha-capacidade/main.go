package main

import "fmt"

func main() {
	s := make([]int, 2, 5)
	t := append(s, 1)
	u := append(s, 2)
	fmt.Println(len(s), cap(s), t, u)
}
