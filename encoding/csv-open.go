/***

Open and read from a CSV file.

***/

package main

import (
	"encoding/csv"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("mention a csv file to open");
		return
	}

	fs, err := os.Open(os.Args[1]); if err != nil {
		fmt.Println("file not found");
		return
	}

	r := csv.NewReader(fs);

	for {
		record, err := r.Read(); if err == io.EOF {break}
		if err != nil {
			fmt.Println(err);
			break
		}
		fmt.Println(record);
	}
}


