
package main

import (
	"fmt"
	"time"
)

func main() {

	sendChan := make(chan int);
	receiveChan := make(chan int);
	go processChannel(sendChan, receiveChan);
	time.Sleep(10e9);

}

func processChannel(in <-chan int, out chan<- int) {
	for inValue := range in {
		result := 0 // Do something, in this case left as nothing.
		out <- result;
	}
}
