package main

import "fmt"

func capitaliza(s string) string {
	if s == "" {
		return s
	}

	b := []byte(s)

	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] -= 'a' - 'A'
	}
	return string(b)
}

func main() {
	a := "luis"
	fmt.Println(capitaliza(a))
	b := "jonice"
	fmt.Println(capitaliza(b))
	c := ""
	fmt.Printf("%q\n", capitaliza(c))
	d := "Gopher"
	fmt.Println(capitaliza(d))
	e := "1abc"
	fmt.Println(capitaliza(e))

}
