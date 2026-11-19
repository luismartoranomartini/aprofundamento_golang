package main

import "fmt"

func main() {

	// a := []int{1, 2, 3, 4, 5}
	// b := []int{9}
	// // a[3] = b[0]
	// //
	// // fmt.Println(a)
	//
	// for i := range a {
	// 	if i == 3 {
	// 		a[i] = b[0]
	// 	}
	// }
	// fmt.Println(a)

	s := []rune("Luís Martini")

	for i, r := range s {
		if r == 'í' {
			s[i] = 'w'
		}
		fmt.Println(string(s[i]))
	}

}
