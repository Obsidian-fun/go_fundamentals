
package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

func CheckError(err error){
	if err != nil {
		log.Fatalf("Get: %v", err);
	}
}


func main() {
	url := "https://golang.org";

	resp, err := http.Get(url);
	CheckError(err);
	defer resp.Body.Close();


	data, err := io.ReadAll(resp.Body);
	CheckError(err);

	fmt.Printf("%q",string(data));

}



