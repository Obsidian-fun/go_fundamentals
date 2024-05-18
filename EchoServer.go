package main

import (
	"os"
	"log"
	"net"
)

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing");
	conn, err := net.DialTCP("tcp", nil, addr);

}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address");
	}

// ResolveTCPAddr returns a TCP endpoint
	addr, err := net.ResolveTCPAddr("tcp",os.Args[1]);
	if err != nil {
		log.Fatalln("Invalid addres: ",os.Args[1], err);
	}

// Listen on that endpoint
	listener, err := net.ListenTCP("tcp", addr);
	if err != nil {
		log.Fatalln("Listener: ",os.Args[1], err);
	}
	log.Println("Listening on ", addr);

// Abillity to create several non-blocking connections to endpoint
	go createConn(addr);

// Accept client side conections to endpoint
	conn, err := listener.AcceptTCP();
	if err != nil {
		log.Fatalln("<- Accept : ", os.Args[1], err);
	}

// Send a response
	handleConn(conn);

}


