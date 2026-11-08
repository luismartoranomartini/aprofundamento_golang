package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Contador struct{ Bytes int }

func (c *Contador) Write(p []byte) (int, error) {
	c.Bytes += len(p)
	return len(p), nil
}

func main() {
	fmt.Fprintln(os.Stdout, "para a tela")

	var sb strings.Builder
	fmt.Fprintf(&sb, "n=%d", 5)
	fmt.Println(sb.String())

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "dentro do buffer")
	fmt.Println(buf.String())

	var c Contador
	fmt.Fprintf(&c, "olá %d", 7)
	fmt.Println(c.Bytes)
}
