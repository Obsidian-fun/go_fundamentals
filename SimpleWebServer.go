
package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServerMux();

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello!");
	});

	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *Request){
		fmt.Fprintf(w,"Goodbye.");
	});

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *Request){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w,"An Error occured [501]");
	});

}
