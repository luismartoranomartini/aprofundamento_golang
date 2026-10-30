package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}
	s := nums[1:4]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s[0] = 99
	fmt.Println(nums)
}
