package main

import "fmt"

func main() {
	comBool := make(map[int]bool)
	comBool[7] = true
	comBool[9] = false

	_, ok9 := comBool[9]
	fmt.Println("bool: comBool[9] =", comBool[9], "| chave 9 existe?", ok9, "| len:", len(comBool))

	comVazio := make(map[int]struct{})
	comVazio[7] = struct{}{}

	_, ok7 := comVazio[7]
	_, ok9v := comVazio[9]
	fmt.Println("struct{}: 7 existe?", ok7, "| 9 existe?", ok9v, "| len:", len(comVazio))
}
