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

	const address = "localhost:8000";

	http.HandleFunc("/", func(w http.ResponseWriter, r *Request) {
		time := time.Second*10 * time.Duration(rand.Intn(10));
	});

}


