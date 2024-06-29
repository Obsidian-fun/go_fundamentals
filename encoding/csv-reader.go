
package main

import (
	"fmt"
	"strings"
	"io"
	"encoding/csv"
)

func main() {
	r := csv.NewReader(strings.NewReader("a,b,c\ne,f,g\n1,2,3"));

	for {
		r, err := r.Read(); if err == io.EOF {break};

		if err != nil {
			fmt.Println("error: ", err);
			break;
		}
		fmt.Println(r);
	}
}
