package main

import "fmt"

func descreve(v any) string {
	switch t := v.(type) {
	case int:
		return fmt.Sprint("int ", t+1)
	case string:
		return fmt.Sprint("string ", len(t))
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("outro %T", t)
	}
}

func main() {
	fmt.Println(descreve(41))
	fmt.Println(descreve("go"))
	fmt.Println(descreve(nil))
	fmt.Println(descreve(2.5))
}
