
package main

import (
	"fmt"
	"net/http"
)

func main() {
	if len(os.Args) != 2 {
		log.Fataln("usage: ./ServeFile.go [Directory]");
	}

// Checking if directory and valid path, os.Stat returns file info.
	s, err := os.Stat(os.Args[1]);
	if err != nil && !s.isDir() {
		err = errors.New("not a directory");
	}

	if err != nil {
		log.Fatalln("invalid path:", err);
	}

// os.Args[1] is now root directory
	http.Handle("/", http.FileServer(http.Dir(os.Args[1])));

	if err := http.ListenAndServe(":3000",s); err != nil {
		log.Fatal(err);
	}

}
