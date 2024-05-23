/***
Handling file upload, with a small form to send the request from a browser
***/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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

	mux := http.NewServerMux();

	mux.HandleFunc(endpoint, func(w http.ResponseWriter,r *http.Request){
		if r.Method == "GET" {
			

	}



