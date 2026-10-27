package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = "50"

	t := reflect.TypeOf(x)
	fmt.Println(t)

	v := reflect.ValueOf(x)
	if v.CanInt() {
		xInt := v.Int()
		xInt = xInt + 1
		fmt.Println(xInt)
	} else {
		fmt.Println("x não é um valor inteiro")
	}

	fmt.Println("----------------")

	var y any = 100
	v2 := reflect.ValueOf(&y).Elem()
	v2.Set(reflect.ValueOf(1000))
	fmt.Println(y, v2.Elem().Int())

}
