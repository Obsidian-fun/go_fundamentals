

package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("please specify a file to write");
	}

	file, err := os.OpenFile(os.Args[1], os.O_WRONLY | os.O_CREATE, 0666);
	if err != nil {
		fmt.Println("cannot create file, %s",os.Args[1]);
	}

	defer file.Close();

	writer := bufio.NewWriter(file);

	writer.WriteString("Jet fuels can't melt steal beams\n");
	writer.Flush();
}

