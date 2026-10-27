package main

import "fmt"

func main() {
	listaToda := []int{2, 10, 9, 4, 5, 8, 1, 3, 6, 7}
	listaToda = append(listaToda, 1)

	fmt.Printf("%T", listaToda)
	fmt.Println(listaToda)
}
