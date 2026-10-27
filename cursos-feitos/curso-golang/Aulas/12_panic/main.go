package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("f:/proprio.pdf")
	

	if err != nil {
		panic(err)
	}


	fmt.Println(err)
}