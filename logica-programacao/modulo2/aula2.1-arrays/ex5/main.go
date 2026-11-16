package main

import "fmt"

func main() {
	var m [2][3]int
	m[1][2] = 5
	fmt.Println(m, len(m), len(m[0]))

	for i := 0; i < len(m); i++ {
		fmt.Print("[ ")
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%d  ", m[i][j])
		}
		fmt.Println(" ]")
	}
}
