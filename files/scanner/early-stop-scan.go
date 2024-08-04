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

// onComma detects comma ",", and the string "STOP"
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.IndexByte(data, ',');
		if i == -1 {
			if !atEOF {
				return 0, nil, nil
			}
			return 0, data, bufio.ErrFinalToken
		}
		if string(data[:i]) == "STOP" {
			return i+1, nil, bufio.ErrFinalToken
		}
		return i+1, data[:i], nil
	}

	input := "1,2,3,4,STOP,5,6,7"

	scanner := bufio.NewScanner(strings.NewReader(input));
	scanner.Split(onComma);

	for scanner.Scan() {
		fmt.Println(scanner.Text());
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err);
	}
}

