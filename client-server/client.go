
package main

import (
	"fmt"
	"os"
	"bufio"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp","localhost:8000");
	if err != nil {
		fmt.Println("error dialing");
		return
	}

	input := bufio.NewReader(os.Stdin);
	fmt.Println("What would you like to be called?");
	clientName, _ := input.ReadString('\n');
	fmt.Printf("Welcome %s \nFeel free to chat!\n", clientName);

	trimmedClient := strings.Trim(clientName,"\n")

	for {

		message, _ := input.ReadString('\n');
		trimmedMessage := strings.Trim(message,"\n");

		if strings.ToLower(trimmedMessage) == "q" {
			fmt.Println("Disconnected");
			return;
		}

		conn.Write([]byte(trimmedClient + ":" + trimmedMessage));
	}
	defer conn.Close();

}

