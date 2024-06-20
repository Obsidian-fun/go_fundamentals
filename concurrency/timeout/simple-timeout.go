/***


***/

package main

import (
//	"fmt"
	"time"
)

func main() {

	ch := make(chan int);
	timeout := make(chan bool, 1);

	go func(){
		time.Sleep(1e9);
		timeout <- true;
	}()

	select {
		case <-ch:
		case <-timeout:
			break;
	}
}

