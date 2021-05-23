package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(bs []byte) (int, error) {
	n, err := r.r.Read(bs)
	for i, b := range bs {
		if b >= 'A' && b <= 'Z' {
			bs[i] = 'A' + ((b - 'A' + 13) % 26)
		} else if b >= 'a' && b <= 'z' {
			bs[i] = 'a' + ((b - 'a' + 13) % 26)
		} else {
			bs[i] = b
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
