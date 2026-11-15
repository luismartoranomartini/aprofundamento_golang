package main

import "fmt"

func main() {
	goto dentro
	for i := 0; i < 2; i++ {
	dentro:
		fmt.Println(i)
	}

}
