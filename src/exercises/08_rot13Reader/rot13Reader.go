package rot13Reader

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (n int, err error) {
	n,err = rot13.r.Read(b)
	for i, _ := range b {
		if b[i] >= 'A' && b[i] < 'N' || b[i] >= 'a' && b[i] < 'n' {
			b[i] += 13
		} else if b[i] >= 'N' && b[i] <= 'Z' || b[i] >= 'n' && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return n, err
}

func Crack() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	io.Copy(os.Stdout, rot13Reader{s})
}