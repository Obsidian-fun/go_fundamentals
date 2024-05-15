/***
Implementation of 'cat' in bash,
Usage: ./ReadFile.go <<File Path>>

***/

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("mention a file pathname to read");
		return
	}

	data, err := os.ReadFile(os.Args[1]);
	if err != nil {
		fmt.Println("file not found");
	}

	os.Stdout.Write(data);
}
