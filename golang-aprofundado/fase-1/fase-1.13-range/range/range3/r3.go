package main

import "fmt"

func main() {
	for i, r := range "aç" {
		fmt.Println(i, r, string(r))
	}
	for i, r := range "a\xffb" {
		fmt.Printf("%d %U\n", i, r)
	}
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
