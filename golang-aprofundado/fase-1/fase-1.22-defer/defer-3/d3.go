package main

import "fmt"

func dobra() (r int) {
	defer func() { r *= 2 }()
	return 5
}

func semNome() int {
	r := 5
	defer func() { r *= 2 }()
	return r
}

func laco() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("laco", i)
	}
	fmt.Println("saindo do laco")
}

func main() {
	fmt.Println(dobra())
	fmt.Println(semNome())
	laco()
}
