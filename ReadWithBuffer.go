
package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	if len(os.Args[1]) > 2 {
		log.Fatalln("provide a file to read");
	}

	f, err := os.Open(os.Args[1]);
	if err != nil {
		log.Println("Error: ", err);
	}

	buf := make([]byte, 16);

	for n:=0; err == nil {
		n, err = f.Read(buf);
		if err == nil {
			fmt.Print(string(buf[:n]));
		}
	}
	if err != nil || err != io.EOF {
		log.Println("\n\nerr: ", err);
	}

}



