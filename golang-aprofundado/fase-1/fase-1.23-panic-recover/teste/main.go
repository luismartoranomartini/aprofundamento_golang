package main

import "fmt"

func g() (s string) {
	defer func() {
		r := recover()
		s = fmt.Sprint("rec:", r)
	}()
	defer fmt.Println("d1")
	fmt.Println("antes")
	panic("boom")
	fmt.Println("depois")
	return "ok"
}

func main() {
	fmt.Println(g())
	fmt.Println("fim")
}
