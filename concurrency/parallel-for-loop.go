/***
Executing a parallel for loop with channels for coordination

***/

package main

import (
	"fmt"
	"time"
)

func doSomething(base float64) float64{
	var padding = 10.0 ;
	base = base + padding;
	return base;
}

func main() {

	data := []float64{3.14,2.3434,323.123,65.5644}
	res := make([]float64,4);

	for i, v := range data {
		go func(i int, v float64) {
			// do something...for example,
			res[i] = doSomething(v);
			fmt.Println(res[i]);
		}(i,v)
	}

	// 0.5 seconds
	time.Sleep(0.5 * 1e9);
}





