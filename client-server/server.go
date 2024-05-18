
package main

import (
	"fmt"
	"log"
	"time"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage: ./server [address:port]");
	}

	addr, err := net.ResolveTCPAddr("tcp", os.Args[1]);
	if err != nil {
		log.Println("invalid address", addr, err);
	}

	listener, err := net.ListenTCP("tcp", os.Args[1]);
	if err != nil {
		log.Println("Listener error : ",os.Args[1], err);
	}




}
