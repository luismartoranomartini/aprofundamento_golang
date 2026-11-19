package main

import "fmt"

func main() {
	t := make([][]int, 4)
	for i := range t {
		t[i] = make([]int, i+1)
		for j := range t[i] {
			t[i][j] = j + 1
		}
	}
	for i, linha := range t {
		fmt.Println(i, linha, len(linha))
	}
}
