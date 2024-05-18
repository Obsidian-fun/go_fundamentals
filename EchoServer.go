package main

import (
	"log"
	"net"
)

func createConn(addr *net.TCPAddr) {


}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Please specify an address");
	}

	addr, err := net.ResolveTCPAddr("tcp",os.Args[1];
	if err != nil {
		log.Fatalln("Invalid addres: ",os.Args[1], err);
	}

	listener, err := net.ListenTCP("tcp", addr);
	if err != nil {
		log.Fatalln("Listener: ",os.Args[1], err);
	}
	log.Println("Listening on ", addr);

	go createConn(addr);

	conn, err := listener.AcceptTCP();
	if err != nil {
		log.Fatalln("<- Accept : ", os.Args[1], err);
	}

	handleConn(conn);

}


