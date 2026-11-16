package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(len(s), cap(s))

	var t []int
	for i := 0; i <= 5; i++ {
		t = append(t, i)
	}
	fmt.Println(len(t), cap(t))
}
