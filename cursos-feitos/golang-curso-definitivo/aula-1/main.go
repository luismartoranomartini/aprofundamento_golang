package main

import "fmt"

func main() {
	// arr := [5]int{}
	sls := []int{}

	// fmt.Println("Array: ", len(arr), cap(arr))
	fmt.Println("Slice: ", len(sls), cap(sls))

	// arr[0] = 10
	sls[0] = 10

	// fmt.Println("Array: ", arr, len(arr), cap(arr))
	fmt.Println("Slice: ", len(sls), cap(sls))
}
