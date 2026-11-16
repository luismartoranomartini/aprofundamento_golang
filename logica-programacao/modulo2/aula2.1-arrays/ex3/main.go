package main

import "fmt"

func main() {
	b := [3]int{10, 20, 30}
	fmt.Println(b == [3]int{10, 20, 30})
	fmt.Println(b == [3]int{10, 20, 31})
}
