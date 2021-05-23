package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(bs []byte) (int, error) {
	for i := range bs {
		bs[i] = 'A'
	}
	return len(bs), nil
}

func main() {
	reader.Validate(MyReader{})
}
