
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





