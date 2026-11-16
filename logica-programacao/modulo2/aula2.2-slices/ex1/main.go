package main

import "fmt"

func main() {
	var a []int
	b := []int{}
	c := []int{10, 20, 30}
	d := make([]int, 3)
	e := make([]int, 2, 5)
	arr := [4]int{1, 2, 3, 4}
	f := arr[:]
	fmt.Println(a, b, c, d, e, f)
}
