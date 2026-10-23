package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	n := copy(dst, src)
	fmt.Println(n, dst)

	clone := make([]int, len(src))
	copy(clone, src)
	clone[0] = 100
	fmt.Println(src, clone)

	a := []int{1, 2, 3, 4}
	b := a[:2:2]
	b = append(b, 99)
	fmt.Println(a, b)

	var n1 []int
	n2 := []int{}
	fmt.Println(n1 == nil, n2 == nil, len(n1), len(n2))

	j1, _ := json.Marshal(n1)
	j2, _ := json.Marshal(n2)
	fmt.Println(string(j1), string(j2))
}
