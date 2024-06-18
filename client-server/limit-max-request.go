
package main

import (
	"fmt"
	"time"
)

const (
	AvailableMemory = 10 << 20; // bit shift right, 10 MB
	AverageMemoryPerRequest = 10 << 10; // 10 KB
	MAXREQS = AvailableMemory/AverageMemoryPerRequest; // 1000 requests
)

// semaphore
var sem = make(chan int, MAXREQS);

type Request struct {
	a, b int;
	replyc chan int;
}


func process(request *Request){
	// Do something..
	fmt.Println("Request processed!");
}

func handle(request *Request){
	process(request);
	// signal done: enable next request to start
	// by making 1 empty place in buffer..
	<-sem;
}

func Server(queue chan *Request) {
	fmt.Println("Server Starting...");
	for {
		sem <- 1;
		// Blocks after request number 1000
		// wait until there is capacity to request
		request := <-queue;
		go handle(request);
	}
}

func main() {
	queue := make(chan *Request);
	go Server(queue);
	time.Sleep(5e9);
}



