package main

import "fmt"

func seguro() (res string) {
	defer func() {
		if r := recover(); r != nil {
			res = fmt.Sprint("recuperado: ", r)
		}
	}()
	panic("deu ruim")
}

func main() {
	fmt.Println(seguro())
	fmt.Println("main continua")
}
