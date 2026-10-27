package main

import "fmt"

func main() {
	tabuada := 5
	i := 1

	for i <= 10 {
		fmt.Println(tabuada * i)
		i++
	}
}