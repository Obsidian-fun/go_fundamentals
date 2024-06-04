
package main

import (
	"fmt"
	"bytes"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("mention a file to read");
	}

	var text, err = os.Open(os.Args[1]);
	if err != nil {
		fmt.Println("file not found\n", err);
		return
	}

	b := bytes.NewBuffer(make([]byte, 26));

	defer text.Close();
	for i:=0; ; i++ {
		b.Reset();
		n := text.Read(b);
		fmt.Printf("%s",string(n[:i]));
	}
}


