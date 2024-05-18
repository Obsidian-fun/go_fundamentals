
package main

import (
	"log"
	"bufio"
	"strings"
	"time"
	"net"
	"os"
)

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
		log.Fatalln("usage: ./server [ip_address:port]");
	}

	addr, err := net.ResolveTCPAddr("tcp", os.Args[1]);
	if err != nil {
		log.Println("invalid address", addr, err);
	}

	listener, err := net.ListenTCP("tcp", addr);
	if err != nil {
		log.Println("Listener error : ",os.Args[1], err);
	}

	for {
		time.Sleep(time.Millisecond * 100);
		conn, err := listener.AcceptTCP();
		if err != nil {
			log.Println("<-Listener : ", os.Args[1], err);
		}
		go handleConn(conn);
	}
}
