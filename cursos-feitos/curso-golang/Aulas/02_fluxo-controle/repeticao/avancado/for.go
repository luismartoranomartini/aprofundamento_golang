package main

import (
	"fmt"
	// "time"
)

func main() {

	// Como se fosse o while
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...\n")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Os números Pares são %d\n", i)
		} else {
			fmt.Printf("Os números Ímpares são %d\n", i)

		}
	}

	// for {
	// 	// laço infinito
	// 	fmt.Println("Para sempre")
	// 	time.Sleep(time.Second)
	// }

	tabuada := 5
	a := 1

	for a <= 10 {
		fmt.Println(tabuada * a)
		a++
	}

}
