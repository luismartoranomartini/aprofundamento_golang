package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println("canal é nil?", ch == nil)
	close(ch)
}
