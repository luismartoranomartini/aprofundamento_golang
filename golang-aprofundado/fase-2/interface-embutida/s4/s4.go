package main

import (
	"bytes"
	"fmt"
	"io"
)

type LeitorEscritor interface {
	io.Reader
	io.Writer
}

func usa(rw LeitorEscritor) {
	fmt.Fprint(rw, "olá")
	b := make([]byte, 10)
	n, _ := rw.Read(b)
	fmt.Println(string(b[:n]))
}

func main() {
	var buf bytes.Buffer
	usa(&buf)
	fmt.Printf("%T\n", &buf)
}
