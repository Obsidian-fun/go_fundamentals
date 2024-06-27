

package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://google.com/",
	"https://golang.org/",
	"https://blog.golang.org/",
}

func main() {

	for _ , url := range urls {
		resp, err := http.Head(url);
		if err != nil {
			fmt.Println("error fetching ", err);
		}

		fmt.Println(url+ ":"+ string(resp.Status));
	}


}





