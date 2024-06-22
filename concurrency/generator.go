/***

Implementing a lazy generaotor in go.
Generator is a function that returns the next value in a sequence when it is called.

***/

package main

import (
	"fmt"
)

var resume chan int;

func integers() chan int {
	yield := make(chan int);
	count := 0;
	go func() {
		for {
			yield <- count;
			count++;
		}
	}()
	return yield;
}

func generateInteger() int {
	return <-resume;
}

func main() {
	resume = integers();
	fmt.Println(generateInteger());
	fmt.Println(generateInteger());
	fmt.Println(generateInteger());
}


