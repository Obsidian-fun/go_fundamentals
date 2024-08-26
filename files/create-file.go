
package main

import (
	"os"
	"log"
)

func main() {
	ParentFile := "Company/";
	ChildFile := "A random directory";

	err := os.MkdirAll(ParentFile+ChildFile, 0750); 
	if err != nil && !os.IsExist(err) {
		log.Fatal(err);
	}
}


