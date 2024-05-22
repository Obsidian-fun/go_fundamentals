
package main

import (
	"fmt"
	"log"
	"net/http"
)

type CustomHandler int

// Counter to keep a measure of endpoint calls to /custom
func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"%d\n",*c);
		*c++;
}

func main() {
	mux := http.NewServeMux();

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello!\n");
	});

	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Goodbye.\n");
	});

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w,"An Error occured.\n");
	});

	mux.Handle("/custom", new(CustomHandler));

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err);
	}
}



