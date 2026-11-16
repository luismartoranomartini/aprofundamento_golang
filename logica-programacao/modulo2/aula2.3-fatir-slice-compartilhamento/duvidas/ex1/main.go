package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 2, 7)
	arr := [6]int{}
	c := arr[2:4]
	d := arr[2:4:5]
	fmt.Println(cap(a), cap(b), cap(c), cap(d))
}
