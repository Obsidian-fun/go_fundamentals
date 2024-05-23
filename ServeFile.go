
package main

import (
	"fmt"
	"net/http"
)

func main() {
	if len(os.Args) != 2 {
		log.Fataln("usage: ./ServeFile.go [Directory]");
	}

	s, err := os.Stat(os.Args[1]);
