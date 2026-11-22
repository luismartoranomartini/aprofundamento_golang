package main

import "fmt"

func main() {
	presentes := make(map[string]bool)

	presentes["Ana"] = true
	presentes["Bia"] = true
	presentes["Ana"] = true

	fmt.Println(len(presentes))

	fmt.Println(presentes["Ana"])
	fmt.Println(presentes["Caio"])

	delete(presentes, "Ana")
	fmt.Println(presentes["Ana"], len(presentes))
}
