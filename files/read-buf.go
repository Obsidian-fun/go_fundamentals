
package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("specify a file to read");
	}

	f, err := os.Open(args[1]);
	if err != nil {
		log.Println("file not found", err);
	}

	defer f.Close();

	var b = make([]byte, 16);
	for i:=0; ;i++ {
		n, err := f.Read(b); if err == nil {
			fmt.Printf("%s",string(b[:i]));
		}
	}
	if err != nil || err != io.EOF {
		fmt.Println("\n\n error reading from file\n");
	}

}
