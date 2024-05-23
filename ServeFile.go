
package main

import (
	"net/http"
	"log"
	"os"
	"errors"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("usage: ./ServeFile.go [Directory]");
	}

// Checking if directory and valid path, os.Stat returns file info.
	s, err := os.Stat(os.Args[1]);
	if err != nil && !s.IsDir() {
		err = errors.New("not a directory");
	}

	if err != nil {
		log.Fatalln("invalid path:", err);
	}

// os.Args[1] is now root directory
	http.Handle("/", http.FileServer(http.Dir(os.Args[1])));

	if err := http.ListenAndServe(":3000",nil); err != nil {
		log.Fatal(err);
	}

}
