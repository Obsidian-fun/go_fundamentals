
package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

type logTripper struct {
	http.RoundTripper;
}

func (l logTripper) logger(r *http.Request) (*http.Response, error)(
	log.Println(r.URL);
	r.Header.Set("X-Log-Time", time.Now().String());

}

func main() {

}
