package main
import "fmt"

func main() {

	for numBase := 1; numBase <= 10; numBase++{
		for mult := 1; mult <=10; mult++{
			fmt.Println(numBase, " X ", mult, " = ", numBase*mult)

		}
	}

}
