/***
Here, the ch channel keeps a track of the computation, if the computation takes longer
than 1 second, the timeout channel halts the program.

***/

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10);
	timeout := make(chan bool, 1);

	go func(){

		for i:=0; i<10;i++ {
			time.Sleep(time.Millisecond*100);
			ch<-i;
		}
		time.Sleep(1e9);
		timeout <- true;
	}()

	for {
		select {
			case value:=<-ch:
				fmt.Println(value);

			case end:=<-timeout:
				if end == true {
					fmt.Println("Finito");
					return;
				}
		}
	}
}

