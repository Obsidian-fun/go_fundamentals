/***

copy-file.go: Usage, ./copy-file.go [SOURCE FILE]

***/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("please specify a file to be copied");
	}


	fmt.Println("Copying...");
	CopyFile(os.Arg[1], "copied.txt");
	fmt.Println("Copy Done!");
}

func CopyFile(destination string, source string){

	src, err := os.Open(source);
	if err != nil {
		fmt.Println("file not found in path");
		return
	}
	defer f.Close();

	dst, err := os.Open(source);
	if err != nil {
		fmt.Println("file not found in path");
		return
	}
	defer f.Close();

	io.Copy(destination, source);
	return
}


