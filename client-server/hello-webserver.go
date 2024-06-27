
package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello "+ r.URL.Path[1:] + "\n");
}

func main() {
	r := http.NewServeMux();

	r.HandleFunc("/obsidian", Hello);

	http.ListenAndServe(":8000",r);
}

