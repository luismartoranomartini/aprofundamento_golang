package main

import "fmt"

func main() {
	goto laco
	fmt.Println("estalinha é pulada")
laco:
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}
