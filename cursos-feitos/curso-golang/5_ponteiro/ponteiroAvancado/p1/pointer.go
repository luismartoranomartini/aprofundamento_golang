package main

import "fmt"

func double(a *int) {
	*a = *a * 2
}

// func double(n int) {
// 	n = n * 2
// 	fmt.Println(n)
// }

func main() {
	a := 10
	fmt.Println(a) // 10
	double(&a)     // 20
	//encerrou a função double
	fmt.Println(a) // 10
}

// func main() {
// 	valor := 10
// 	double(valor)
// }
