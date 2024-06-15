/***
Executing a parallel for loop with channels for coordination

***/

package main

import (
	"fmt"
)

func doSomething(base float64) float64{
	var padding = 10.0 ;
	base = base + padding;
	return base;
}

func main() {

	data := []float64{3.14,2.3434,323.123,65.5644}
	res := make([]float64,4);

	type Empty interface{};
	var empty Empty;

// Creating a semaphore (buffered channel size of data and Empty datatype, 
	sem := make(chan Empty,4);

	for i, v := range data {
		go func(i int, v float64) {
			// do something...for example,
			res[i] = doSomething(v);
			fmt.Println(res[i]);
			sem <- empty;
		}(i,v)
	}

// make main wait for goroutines to finish
	for i:=0; i<4; i++ {
		<-sem;
	}
}





