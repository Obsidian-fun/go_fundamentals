/***
ALSO IDIOMATIC,
Instead of passing the channel into the function, let the function generate the channels
function plays the role of a "factory"
***/

package main

import (
	"fmt"
	"time"
)

func main() {

	stream := pump();
	go suck(stream);  // could be "suck(pump())"

	time.Sleep(1e5);
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
	for {
		fmt.Println(<-ch);
	}
}
