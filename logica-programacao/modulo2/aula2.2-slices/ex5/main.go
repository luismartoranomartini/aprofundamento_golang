package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]
	fmt.Println(s, len(s), cap(s))
	s[0] = 99
	fmt.Println(arr)
	arr[3] = 77
	fmt.Println(s)
}
