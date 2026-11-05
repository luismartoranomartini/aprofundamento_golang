package main

import "fmt"

func bissexto(ano int) bool {
	if ano%400 == 0 {
		return true
	}
	if ano%100 == 0 {
		return false
	}
	if ano%4 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(bissexto(2000))
	fmt.Println(bissexto(1900))
	fmt.Println(bissexto(2024))
	fmt.Println(bissexto(2023))
}
