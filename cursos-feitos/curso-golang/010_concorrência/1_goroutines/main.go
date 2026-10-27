package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("bom dia!")
	escrever("programando em Go")
}

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
