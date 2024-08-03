/***
Source: go docs, "pkg.go.dev/bufio"
***/

package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	input := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input));

	scanner.Split(bufio.ScanWords);
	count := 0; // count the words

	for scanner.Scan() {	
		fmt.Println(scanner.Text());
		count++;
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error reading input: ", err);
	}

	fmt.Printf("Words counted: %d\n", count);


}
