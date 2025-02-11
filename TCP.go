package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"log"
	"net"
)

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing");
	conn, err := net.DialTCP("tcp", nil, addr);
	if err != nil {
		log.Fatalln("-> Connection:",err);
	}
	log.Println("Connection to ->",addr);

	r := bufio.NewReader(os.Stdin);
	for {
		fmt.Print("# ");
		msg, err := r.ReadBytes('\n');
		if err != nil {
			log.Println("Message error ", msg, err);
		}

		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connection", err);
			return;
		}
	}
}

func handleConn(conn net.Conn) {
	r := bufio.NewReader(conn);
	time.Sleep(time.Second/2);
	for {
		msg, err := r.ReadString('\n');
		if err != nil {
			log.Println("<- Message Error", msg, err);
		}

		msg = strings.TrimSpace(msg);
		switch msg {
			case `\q`:
				log.Println("Exiting....");
				if err := conn.Close(); err != nil {
					log.Println("Error closing connection");
				}
				time.Sleep(time.Second/2);
				return;

			case `\x`:
				log.Println("Special message `\\x` received!");

			default:
				log.Println("Message received");
		}
	}
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


