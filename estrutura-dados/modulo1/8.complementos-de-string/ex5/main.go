package main

import "fmt"

func main() {
	a := "go"
	b := a + "lang"

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	s := ""
	for i := 0; i < 3; i++ {
		s += "x"
		fmt.Printf("%q ", s)
	}

	buf := []byte{}
	for i := 0; i < 100000; i++ {
		buf = append(buf, 'x')
	}
	r := string(buf)
	fmt.Println(r)

	a = "ab"
	b = a
	a += "c"
	fmt.Println(a, b)
}
