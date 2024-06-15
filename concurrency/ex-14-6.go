/***
Exercise 14.6 from "The way to Go" by Ivo Balbaert.
Program : 2 functions, 
	a) Producer : produces numbers 10, 20, 30.....90 and puts them on a channel. 
  b) Consumer : reads from them and prints them, while main() waits for goroutines to finish.

***/

package main

import "fmt"

const N=9; // size of buffer
type semaphore chan int;

func (s semaphore) producer(n int) {
	var number int;
	for i:=0; i<9; i++ {
		number = 10+(10*i);
		s <-number;
	}
}

func (s semaphore) consumer(n int) {
	var number int;
	for i:=0; i<1; i++ {
		fmt.Printf("%d\n",number);
	}
}

func main() {
	sem := make(semaphore, N);

	go sem.producer(N);
	go sem.consumer(N);

}

