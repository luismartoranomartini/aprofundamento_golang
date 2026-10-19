package main

import "fmt"

func total(nums ...int) int {
	fmt.Printf("%T len=%d nil=%v\n", nums, len(nums), nums == nil)
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

func troca(nums ...int) {
	nums[0] = 99
}

func main() {
	fmt.Println(total(1, 2, 3))
	fmt.Println(total())

	s := []int{10, 20}
	fmt.Println(total(s...))
	troca(s...)
	fmt.Println(s)

	troca(1, 2)
}
