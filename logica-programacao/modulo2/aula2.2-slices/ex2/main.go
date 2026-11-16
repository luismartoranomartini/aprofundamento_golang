package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	c := []int{}
	var n []int
	fmt.Println(len(a), cap(a))
	d := make([]int, 2, 5)
	fmt.Println(len(d), cap(d))
	fmt.Println(len(c), cap(c), c == nil)
	fmt.Println(len(n), cap(n), n == nil)
	fmt.Printf("%v %v\n", c, n)

}
