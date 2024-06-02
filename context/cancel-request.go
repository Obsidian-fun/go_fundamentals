
/*** Cancel an HTTP request client side,
In this program, a client will attempt to connect to a server, and if it fails to do so within deadline
(2 seconds in this example), it will close the connection.

For demonstration, the server has been set to sleep for 5 seconds before it serves the root file.
As it is greater than the dealine set, the connection will be closed after 2 seconds.
***/

package main

import (
	"fmt"
	"log"
	"time"
	"context"
	"net/http"
)

func main() {

	address := "localhost:8000";

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second*5);
		fmt.Fprint(w,"Hello there...");
	});

	req,err := http.NewRequest("GET","http://"+address, nil);
	if err != nil {
		fmt.Println(http.StatusNotFound,err);
	}
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*2);
	defer cancel();

	go func() {
		if err := http.ListenAndServe(address,nil); err != nil {
			log.Fatal(err);
		}
	}()

	if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
		log.Fatalln(err);
	}

}
