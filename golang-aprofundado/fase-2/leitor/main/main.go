package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Leitor struct {
	dados string
	pos   int
}

func NovoLeitor(s string) *Leitor {
	return &Leitor{dados: s}
}

func (l *Leitor) Read(p []byte) (int, error) {
	if l.pos >= len(l.dados) {
		return 0, io.EOF
	}

	n := copy(p, l.dados[l.pos:])
	l.pos += n
	return n, nil

}

func main() {
	texto := "um\ndois\ntres\n"

	l := NovoLeitor(texto)
	buf := make([]byte, 4)
	for {
		n, err := l.Read(buf)
		fmt.Printf("n=%d %q err=%v\n", n, buf[:n], err)
		if err != nil {
			break
		}
	}

	n, err := io.Copy(os.Stdout, NovoLeitor(texto))
	fmt.Println("copiados:", n, err)

	sc := bufio.NewScanner(NovoLeitor(texto))
	for sc.Scan() {
		fmt.Println("linha:", sc.Text())
	}
	fmt.Println(sc.Err())
}
