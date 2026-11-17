package main

import "fmt"

func main() {
	src := []int{10, 20, 30, 40, 50}

	dst1 := make([]int, 3)
	n1 := copy(dst1, src)
	fmt.Println("dst maior: n=", n1, "dst =", dst1)

	dst2 := make([]int, 8)
	n2 := copy(dst2, src)
	fmt.Println("dst maior n=", n2, "dst =", dst2)

	dst3 := make([]int, 5)
	copy(dst3, src)
	dst3[0] = 99
	fmt.Println("independente: src =", src, "| dst3 =", dst3)

	dst4 := make([]int, 2)
	copy(dst4, src[3:])
	fmt.Println("pedaço src[3:]:", dst4)
}
