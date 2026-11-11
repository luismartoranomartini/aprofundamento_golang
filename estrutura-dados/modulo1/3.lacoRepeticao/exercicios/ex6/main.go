package main

import "fmt"

func main() {
	n := []byte("12345")

	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	fmt.Println(string(n))
}
