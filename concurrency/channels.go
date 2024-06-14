/***
Channel can be of any datatype including {}, and another channel!

Initialization : var ch1 chan string
ch1 = make(chan string) or, ch1 := make(chan string);

Communication : ch <- int1 (loading the channel)
								int2 = <-ch (loading the variable with channel info)
								<-ch (unload current value and take next value)
								
***/

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string);

	go sendData(ch);
	go getData(ch);

	time.Sleep(1e9); // Sleep for a second before finishing...

}

func sendData(ch chan string) {
// Loading Channels..
	ch <- "Washington";
	ch <- "Los Angeles";
	ch <- "Beijing";
	ch <- "Doha";
}

func getData(ch chan string) {
	var location string;
	fmt.Println("Now visiting.. : \n");
	for {
		location = <-ch;
		fmt.Printf("%s\n", location);
	}
}


