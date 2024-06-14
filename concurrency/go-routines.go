package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Prinln("In main()");

	go longWait();
	go shortWait();
}

func longWait() {
	fmt.Println("Beginning long wait");

}
