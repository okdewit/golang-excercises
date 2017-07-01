package readers

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func Validate() {
	reader.Validate(MyReader{})
}