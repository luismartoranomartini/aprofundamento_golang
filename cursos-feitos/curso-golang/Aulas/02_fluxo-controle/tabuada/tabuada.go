package main

import "fmt"

func main() {

	tabuada := 5
	for i := 1; i <= 10; i++{
		fmt.Printf("%d x %d = %d\n", i, tabuada, tabuada*i)
	}

}