package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	for i := 0; i <= 10; i++{
		go showMsg(strconv.Itoa(i))
	}
	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))

}

func showMsg(msg string) {
	fmt.Println(msg)
}