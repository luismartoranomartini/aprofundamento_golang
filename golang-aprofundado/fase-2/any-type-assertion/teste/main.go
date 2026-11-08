package main

import (
	"errors"
	"fmt"
)

func desc(x any) string {
	switch v := x.(type) {
	case nil:
		return "nil"
	case int, int64:
		return fmt.Sprintf("inteiro %v (%T)", v, v)
	case string:
		return "texto " + v
	case []int:
		return fmt.Sprint("slice ", len(v))
	case error:
		return "erro " + v.Error()
	default:
		return fmt.Sprintf("outro %T", v)
	}
}

func main() {
	var e error
	fmt.Println(desc(7))
	fmt.Println(desc(int64(8)))
	fmt.Println(desc([]int{1, 2}))
	fmt.Println(desc(3.5))
	fmt.Println(desc(e))
	fmt.Println(desc(errors.New("x")))
}
