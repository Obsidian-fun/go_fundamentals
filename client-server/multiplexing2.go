/***
Allowing the user to voluntair 

***/

package main

import "fmt"

type Request struct {

	a, b int;
	replyc chan int;

}

type binOp func (a , b int) int;

func run(op binOp, req *Request){
	req.replyc <- op(req.a, req.b);
}

func server(op binOp, service chan *Request, quit chan bool){
	for {
		select {
			case req := <-service:
					go run(op, req);
			case <-quit:
					return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan Request);

	quit = make(chan bool);

	go server(op, service, quit);
	return service, quit;

}

func main() {



}








