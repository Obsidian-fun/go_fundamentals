/*** 
Query different search engines,
Retrives the result from the fastest one, and cancel all the others..
***/

package main

import (
	"context"
	"time"
	"rand"
	"net/http"
	"log"
	"fmt"
)

func main() {

	// Server side
	const address = "localhost:8000";

	http.HandleFunc("/", func(w http.ResponseWriter, r *Request) {
		duration := time.Second*10 * time.Duration(rand.Intn(10));
		log.Println("wait ",duration);
		time.Sleep(duration);
	});

	go func() {
		if err,_ := http.ListenAndServe(address,nil); err != nil {
			log.Fatal(err);
		}
	]()

	// Client-side,

	ctx, cancel := context.Context



}

