package main

import "fmt"

func f() (n int) {
	n = 1
	defer fmt.Println("a", n)
	defer func() { n += 10 }()
	n = 2
	return n * 3
}

func main() {
	fmt.Println(f())
}
