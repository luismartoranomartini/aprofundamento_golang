package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3, 10)
	n := copy(dst, src)
	fmt.Println(n, dst, len(dst), cap(dst))
}
