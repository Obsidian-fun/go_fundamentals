package main

import (
	"fmt"
	"time"
)

// time.Sleep() operated in nano seconds mo a second is time.Sleep(1e9).

func main() {
	fmt.Println("In main()");

	go longWait();
	go shortWait();

	time.Sleep(10 * 1e9);
	fmt.Println("End of main()");
}

func longWait() {
	fmt.Println("Beginning long wait");
	time.Sleep(5 * 1e9);
	fmt.Println("End of long wait");
}

func shortWait() {
	fmt.Println("Beginning of short wait");
	time.Sleep(2 * 1e9);
	fmt.Println("End of short wait");
}
