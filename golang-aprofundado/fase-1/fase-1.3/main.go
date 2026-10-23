package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 99
	fmt.Println(a, b)

	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:3]
	fmt.Println(s, len(s), cap(s))

	s[0] = 777
	fmt.Println(arr)

	arr1 := [6]int{0, 1, 2, 3, 4, 5}
	s1 := arr1[2:4]
	s1[1] = 100
	fmt.Println(len(s1), cap(s1), arr1)
}
