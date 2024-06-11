
package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader struct {
	src string;
	cur int;
}

func newAlphaReader(src string) * alphaReader {
	return &alphaReader {src:src}
}

// alpha reads only alphabets and reject other special characters
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r;
	}
	return
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src)  {
		return
	}

	x := len(a.src) - a.cur;
	n, bound := 0,0;
	if x > len(p) {
			bound = len(p);
	} else if x <= len(p) {
			bound = x;
	}

	buf := make([]byte, bound);
	for n < bound {
			if char := alpha(a.src[a.cur]); char != 0 {
					buf[n] = char;
			}
			n++;
			a.cur ++;
	}
	copy(p, buf);
	return
}

func main() {
	reader := newAlphaReader("Hello there, breakfast at 8am, it is!");
	p := make([]byte, 4);

	for {
		n, err := reader.Read(p);
		if err == io.EOF {
			break;
		}
		fmt.Print(string(p[:n]));
	}
	fmt.Println();
}





