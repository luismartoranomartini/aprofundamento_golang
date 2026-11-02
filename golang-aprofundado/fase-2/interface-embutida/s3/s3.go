package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("abcdef")
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Printf("n=%d dados=%q err=%v\n", n, buf[:n], err)
		if err == io.EOF {
			break
		}
	}

	n, err := io.Copy(os.Stdout, strings.NewReader("copiado\n"))
	fmt.Println(n, err)
}
