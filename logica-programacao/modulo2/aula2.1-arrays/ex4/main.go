package main

import "fmt"

func dobrar(a [3]int) {
	for i := range a {
		a[i] *= 2
	}
	fmt.Println("dentro", a)
}

func main() {
	b := [3]int{10, 20, 30}
	e := b
	e[0] = 0
	fmt.Println(b, e)
	dobrar(e)
	fmt.Println("fora:", b)

}
