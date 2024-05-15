
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("mention file path");
		return
	}

	root := filepath.Abs


}

