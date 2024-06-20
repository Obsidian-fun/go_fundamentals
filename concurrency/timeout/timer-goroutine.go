
package main

import (
	"fmt"
	"time"
)



func main(){
	tick := time.Tick(1e9);
	boom := time.After(5e9);

	for {
		select {
			case <-tick:
				fmt.Println("Tick...");
			case <-boom:
				fmt.Println("Boom!");
				return
			default:
				fmt.Println();
				time.Sleep(1e8);
		}
	}
}



