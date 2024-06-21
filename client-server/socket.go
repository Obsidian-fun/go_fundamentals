/***

Reading, Writing from a socket

***/

package main

import (
	"fmt"
	"log"
	"net"
	"io"
)

func main() {

	var (
		host = "somerandomexamplesite.com"
		port = "80"
		remote = host + ":" + port
		msg string = "GET / \n"
		data = make([]uint8, 4096)
		read = true
		count = 0
	)

	// create the socket,

	con, err :=  net.Dial("tcp", remote);
	if err != nil {
		log.Fatal(err);
	}
	io.WriteString(con, msg);

	// read the response from the web server,

	for read {
		count, err = con.Read(data);
		read = (err == nil);
		fmt.Printf(string(data[0:count]));
	}
	con.Close();
}

