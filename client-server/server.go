
package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("Starting the server...");
	listener, err := net.Listen("tcp",":8000");
	if err != nil {
		fmt.Println("error : ", err);
		return
	}

	for {
		conn, err := listener.Accept();
		if err != nil {
			fmt.Println("error connecting, ", err);
			return
		}


	}
}


