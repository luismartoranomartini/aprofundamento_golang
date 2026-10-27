package main

import "fmt"

func main() {

	slice1 := []int{5, 1, 2, 3}
	slice2 := []string{"a", "e", "f", "b"}
	newInts := reverse(slice1)
	newString := reverse(slice2)
	fmt.Println(newInts)
	fmt.Println(newString)
	
}

type facil interface {
	int | string
}

func reverse[T facil](slice []T) []T {
	newInts := make([]T, len(slice))

	newIntsLen := len(slice) -1

	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}



