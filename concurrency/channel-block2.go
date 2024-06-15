/***
Continuing with the example, channel will ouput values thanks to the "suck", for about 5 seconds 
before finishing the program.

***/

package main

import (
	"fmt"
	"time"
)

// infinit loop,
func pump(ch chan int) {
	for i:=0; ; i++ {
		ch <- i;
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch);
	}
}

func main() {
	ch := make(chan int);
	go pump(ch);
	go suck(ch);
	fmt.Println(<-ch);

	time.Sleep(5 * 1e9);
}
