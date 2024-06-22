/***

An echo server in go. Make sure to use nc/telnet since curl won't work
(no headers in response = HTTP 0.9, which is blocked in libcurl by default)

***/

package main

import (
	"fmt"
	"net"
	"io"
)

func doServerStuff(conn net.Conn){

	defer conn.Close();

	buf := make([]byte, 1024);

	for {
		_, err := conn.Read(buf[0:]);
		if err == io.EOF {
			fmt.Println("client disconnected");
			return
		}

		fmt.Printf("Received data: %v \n", string(buf));

	}
}

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

		go doServerStuff(conn);
	}
}


