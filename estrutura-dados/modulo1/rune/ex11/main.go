package main

import "fmt"

func main() {
	for _, r := range "a\xffb" {
		fmt.Println(r)
	}
}
