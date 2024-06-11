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
		return
	}

	fmt.Println("Copying...");
	CopyFile("copied.text", os.Args[1]);
	fmt.Println("Copy Done!");
}

func CopyFile(destination string, source string){

	src, err := os.Open(source);
	if err != nil {
		fmt.Println("file not found in path");
		return
	}
	defer src.Close();

	dst, err := os.Open(source);
	if err != nil {
		fmt.Println("file not found in path");
		return
	}
	defer dst.Close();

	io.Copy(dst, src);
	return
}


