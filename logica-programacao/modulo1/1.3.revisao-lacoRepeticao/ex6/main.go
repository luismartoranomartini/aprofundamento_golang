package main

import "fmt"

func main() {
	n, r := 1234, 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	fmt.Println(r)
}
