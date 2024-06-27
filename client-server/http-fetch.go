
package main

import (
	"fmt"
	"net/http"
	"io"
)

func CheckError(err error){
	if err != nil {
		fmt.Println(err);
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



