
package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

func main() {

	if len(os.Args) != 3 {
		log.Println("please specify a file name and text");
		return
	}

	if err := os.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
		fmt.Println("err: ", err);
	}

}

