/***
Handling file upload, with a small form to send the request from a browser
***/

package main

import (
	"fmt"
	"log"
	"io"
	"os"
	"path/filepath"
	"net/http"
)

	const (
		param = "file";
		endpoint = "/uplaod";
		content = `<html lang="en">`+
              `<body>`+
                `<form action="%s" enctype="multipart/form-data" method="POST">`+
                   `<input type="text" name="%s">`+
                   `<input type="submit" value="Upload">`+
               `</form>`+
              `</body>`+
              `</html>`
	)

	func main() {
		mux := http.NewServeMux();

		mux.HandleFunc(endpoint, func(w http.ResponseWriter,r *http.Request){
			if r.Method == "GET" {
				fmt.Fprintf(w, content, endpoint, param);
			}else if r.Method == "POST" {
					http.NotFound(w, r);
			}

			path, err := upload(r);
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError);
					return;
			}
			fmt.Fprintf(w, "Uploaded to %s",path);
		});

		if err := http.ListenAndServe(":8000",mux); err != nil{
				log.Fatal(err);
		}
}

func upload(r *http.Request) (string, error) {
// Get file
	f, h, err := r.FormFile(param);
	if err != nil {
		return "",err;
	}
	defer f.Close();

// File write
	p := filepath.Join(os.TempDir(), h.Filename);
	fw, err := os.OpenFile(p, os.O_WRONLY | os.O_CREATE, 0666);
	if err != nil{
		return "", err;
	}
	defer fw.Close();

// Copy written file to original file, updating it.
	if _ , err := io.Copy(fw, f); err != nil {
		return "",err;
	}

	return p, nil;
}

