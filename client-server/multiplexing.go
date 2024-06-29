
package main

import "fmt"

// Request is a struct with a reply channel embedded inside

type Request struct {
	a       int;
	b       int;
	replyc  chan  int;  //  reply  channel  inside  the  Request
}

// or,
type Reply struct {}

type binOp func(a, b int) int;

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request) {
	for {
		req := <-service; // requests arrive here
		go run(op, req);
	}
}

func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request);
	go server(op, reqChan);
	return reqChan;
}

func main() {

	addr := startServer(func(a, b int) int { return a + b})
	const N=1000000;
	var reqs [N]Request;

	for i:=0; i<N; i++ {
		req := &reqs[i];
		req.a = i;
		req.b = i + N;
		req.replyc = make(chan int);
		addr <- req // addr is a channel of requests
	}

	// checks:
	for i:=N-1; i >=0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at",i);
		} else {
				fmt.Println("Request", i, " is OK!");
		}
	}
	fmt.Println("done");

}





