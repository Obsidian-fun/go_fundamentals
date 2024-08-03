package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin);

	for scanner.Scan() {
		fmt.Println(scanner.Text()); // Println adds the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading standard input:", err);
	} 

}
