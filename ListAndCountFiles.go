/***
Usage:
go build -ldflags "-w -s" ListAndCountFiles.go
./ListAndCountFiles.go <<Name of file in directory>>

***/

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

	// get absolute path from root,
	root, err := filepath.Abs(os.Args[1]);
	if err != nil {
		fmt.Println("cannot get file path");
		return
	}

	fmt.Println("Listing files in root:\n", root);

	var c struct {
		file int;
		dir	 int;
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				c.dir++;
			} else {
				c.file++;
			}

		fmt.Println("-",path);
		return nil;
	})

	fmt.Printf("Total %d files and %d directories",c.file, c.dir);
}

