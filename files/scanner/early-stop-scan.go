/***
Source: go docs, "pkg.go.dev/bufio"
***/

package main

import (
	"fmt"
	"bufio"
	"strings"
	"bytes"
	"os"
)

func main() {
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.IndexByte(data, ',');
		if i == -1 {
			if !atEOF {
				return 0, nil, nil
			}
			return 0, data, bufio.ErrFinalToken
		}



	}


	scanner := bufio.NewScanner();
}



