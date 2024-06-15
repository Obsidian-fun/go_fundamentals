/***
Exercise 14.6 from "The way to Go" by Ivo Balbaert.
Program : 2 functions, 
	a) Producer : produces numbers 10, 20, 30.....90 and puts them on a channel. 
  b) Consumer : reads from them and prints them, while main() waits for goroutines to finish.

***/

package main

import (
	"fmt"
	"time"
)

const N=9; // size of buffer
type semaphore chan int;

func (s semaphore) producer(n int) {
	var number int;
	for i:=0; i<N; i++ {
		number = 10+(10*i);
		s <-number;
	}
}

func (s semaphore) consumer(n int) {
	var number int;
	for i:=0; i<N; i++ {
		number = <-s;
		fmt.Println(number);
	}
}


func main() {
	sem := make(semaphore, N);

	go sem.producer(N);
	go sem.consumer(N);

	time.Sleep(0.1*1e9);

}

