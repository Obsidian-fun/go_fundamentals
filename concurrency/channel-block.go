/***
Without a receiver, channel replies with 0 (first value), since channel is effectively blocked.
***/

package main

import "fmt"

// infinit loop,
func pump(ch chan int) {
	for i:=0; ; i++ {
		ch <- i;
	}
}

func main() {
	ch := make(chan int);
	go pump(ch);
	fmt.Println(<-ch);

}
