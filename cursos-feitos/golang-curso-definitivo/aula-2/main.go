package main

import "fmt"

func main() {

	nums := make([]int, 0, 5)

	fmt.Println(nums, len(nums), cap(nums))

	novosNums := []int{1, 2, 3, 4, 5, 6}

	for _, num := range novosNums {
		nums = append(nums, num)
	}
	fmt.Println(nums)
	fmt.Println(nums, len(nums), cap(nums))
}
