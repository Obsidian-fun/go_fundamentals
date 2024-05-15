
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

	// get absolute path,
	root, err := filepath.Abs(os.Args[1]);
	if err != nil {
		fmt.Println("cannot get file path");
		return
	}

	fmt.Println("Listing files in root", root);

	var c struct {
		file int;
		dir	 int;
	}

	filepath.Walk

}

