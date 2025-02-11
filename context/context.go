/***
Outgoing request to a server with context, to timeout 
Taken from, https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
***/

package main

import (
	"context"
	"net/http"
	"io"
	"os"
	"log"
	"time"
)

func main() {
	// making a request,
	req, err := http.NewRequest("GET",
		"https://www.ardanlabs.com/blog/post/index.xml",
		nil);
	if err != nil {
		log.Println("error making request : ", err);
		return
	}
	// adding a  context,
	ctx, cancel := context.WithTimeout(req.Context(), 10000*time.Millisecond);
	defer cancel()
	// bind the new context into the new request,
	req = req.WithContext(ctx);
	// make the call, Do handles context level timeout,
	resp, err := http.DefaultClient.Do(req);
	if err != nil {
		log.Println("error fetching response, ", err);
		return
	}
	defer resp.Body.Close();
	io.Copy(os.Stdout, resp.Body);

}
