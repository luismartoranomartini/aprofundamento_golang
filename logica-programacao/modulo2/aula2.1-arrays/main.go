package main

import "fmt"

func main() {
	var a [3]int
	b := [3]int{10, 20, 30}
	c := [...]string{"x", "y", "z", "w"}
	d := [5]int{1: 7, 3: 9}
	var e = [3]int{}
	f := [3]int{}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	for i := 0; i < len(b); i++ {
		fmt.Print(i, ":", b[i], " \n")
	}

	for i, v := range b {
		fmt.Print(i, "=", v, " \n")
	}
	fmt.Println(b == [3]int{10, 20, 30})
	fmt.Println(b == [3]int{10, 20, 31})
}
