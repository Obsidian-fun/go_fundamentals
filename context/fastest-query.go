/*** 
Query different search engines,
Retrives the result from the fastest one, and cancel all the others..
***/

package main

import (
	"context"
	"time"
	"net/http"
	"math/rand"
	"sync"
	"log"
)

func main() {

	// Server side
	const address = "localhost:8000";

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Second*10 * time.Duration(rand.Intn(10));
		log.Println("wait ",duration);
		time.Sleep(duration);
	});

	go func() {
		if err := http.ListenAndServe(address,nil); err != nil {
			log.Fatal(err);
		}
	}()

	// Client-side,

	ctx, cancel := context.WithCancel(context.Background());
	ch, o, wg := make(chan int), sync.Once{}, sync.WaitGroup{};

	wg.Add(10);
	for i:=0; i<10; i++ {
		go func(i int) {
			defer wg.Done();
			req,_ := http.NewRequest("GET","http://"+address,nil);
			if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
					log.Println(i, err);
					return
			}
			o.Do(func() {ch <- i})
		}(i)
	}

	log.Println("Recieved", <-ch);
	cancel();
	log.Println("Cancelling");
	wg.Wait();
}

