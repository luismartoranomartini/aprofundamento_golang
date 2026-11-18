package main

import "fmt"

func contem(s []int, v int) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func unicos(s []int) []int {
	var res []int
	for _, v := range s {
		if !contem(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(unicos([]int{3, 1, 3, 2, 1, 2, 6}))
	fmt.Println(unicos([]int{4, 4, 4}))
	fmt.Println(unicos([]int{1, 2, 3}))
	fmt.Println(unicos(nil))
}
