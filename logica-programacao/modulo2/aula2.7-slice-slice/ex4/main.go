package main

import "fmt"

func main() {
	m := [][]int{{1, 2}, {3, 4}}

	m[0] = append(m[0], 9)
	fmt.Println(m)

	m = append(m, []int{7, 7, 7})
	fmt.Println(m, len(m))
}
