
package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func main() {

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error){
		advance, token, err = bufio.ScanWords(data, atEOF);
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token),10,32);
		}
		return
	}

	const input = "12345 378846 21218946823746893456923478634483489567";
	scanner := bufio.NewScanner(strings.NewReader(input));
	
	scanner.Split(split);

	for scanner.Scan() {
		fmt.Printf("%s\n",scanner.Text());
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error scanning: %s", err);
	}
}


