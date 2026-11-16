package main

import "fmt"

func semRetorno(s []int) {
	s = append(s, 1)
}

func comRetorno(s []int) []int {
	return append(s, 1)
}

func main() {
	s := make([]int, 0, 5)

	semRetorno(s)
	fmt.Println("depois de semRetorno:", s, len(s))
	fmt.Println("mas o array recebeu: ", s[:1])

	s = comRetorno(s)
	fmt.Println("depois de comRetorno:", s, len(s))
}
