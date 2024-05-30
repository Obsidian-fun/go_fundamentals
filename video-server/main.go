
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {

	const dirName = "video-server";
	const port = "5555";

	// addHeaders is a custom func, to bypass CORS
	http.Handle("/",addHeaders

}

func addHeaders(h http.Handler) h.http.HandlerFunc




