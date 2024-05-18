package main

import (
	"log"
	"net"
	"bufio"
	"os"
)

func createConn(addr *net.TCPAddr) {
	defer log.Println("-> Closing");

	conn, err := net.DialTCP("tcp", nil, addr);
	if err != nil {
		log.Println("connection error: ", err);
	}
	log.Println("Connection to ->", addr);

	r := bufio.NewReader(os.Stdin);

	for {
		msg, err := r.ReadBytes('\n');
		if err != nil {
			log.Println("message error: ", err);
		}

		if _, err := conn.Write(msg); err != nil {
			log.Println("-> Connection", err);
			return
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("./client [IP_Address:Port]");
	}

	addr, err := net.ResolveTCPAddr("tcp", os.Args[1]);
	if err != nil {
		log.Println(" invalid address: ", os.Args[1], err);
	}

	createConn(addr);
}
