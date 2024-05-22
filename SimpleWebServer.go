
package main

import (
	"fmt"
	"net/http"
)

type CustomHandler int

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"%d",*c);
		*c++;
}

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

	mux.HandleFunc("/custom", new(CustomHandler));

}



