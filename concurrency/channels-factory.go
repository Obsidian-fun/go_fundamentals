/***
ALSO IDIOMATIC,
Instead of passing the channel into the function, let the function generate the channels

***/

package main

import (
	"fmt"
	"time"
)

func main() {

	stream := pump();
	go suck(stream);

	time.Sleep(1e9);
}

func pump() chan int{

	ch:=make(chan int);

	go func() {
		for i:=0; i<10; i++ {
			ch <- i;
		}
	}();

	return ch;
}

func suck(ch chan int) {
	for i:=0; i<10; i++ {
		fmt.Println(<-ch);
	}
}
