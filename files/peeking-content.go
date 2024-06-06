/***
Peeking is the ability to read content without advancing the reader cursor.
Here, under the hood, the peeked data is stored in the buffer. Each reading
operation checks whether there's data in this buffer and if there is any, that
data is returned while removing it from the buffer. This works like a queue
(first in, first out).

This ability allow's us to split the content on a particular condition, such as a
'new line' split to store chunks of lines for example, or a 'space split' to split on words.

The structure that allows an application to achieve this behavior is bufio.Scanner . 
This makes it possible to define what the splitting function is and has the following type:
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

***/

package main

import (
	"fmt"
	"log"
	"bufio"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please specify a path");
	}

	f, err := os.Open(os.Args[1]);
	if err != nil {
		log.Println("file not found ", err);
		return
	}

// Wrapping the reader in a 'buffered' reader,
	r := bufio.NewReader(f);

	var rowCount int;

	for err==nil {
		var b []byte;



	}

}

