
// Cancel an HTTP request client side,

package main

import (
	"fmt"
	"log"
	"time"
	"context"
	"net/http"
)

func main() {

	address := "127.0.0.1:8000";

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second*5);
		fmt.Fprint(w,"Hello there...");
	});

	req,err := http.NewRequest("GET","http://"+address, nil);
	if err != nil {
		fmt.Println(http.StatusNotFound,err);
	}
	fmt.Println(req);
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*2);
	defer cancel();

	go func() {
		err := http.ListenAndServe(address,nil); if err != nil {
			log.Fatal(err);
		}
	}()

	if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
		log.Fatalln(err);
	}

}
