
package main

import (
	"log"
	"time"
	"net/http"
)

type logTripper struct {
	http.RoundTripper;
}

func (l logTripper) logger(r *http.Request) (*http.Response, error){
	log.Println(r.URL);
	r.Header.Set("X-Log-Time", time.Now().String());
	return l.RoundTripper.RoundTrip(r);
}

func main() {
	client := http.Client{Transport : logTripper{http.DefaultTransport}};
	req, err := http.NewRequest("GET","https://google.com/search?q=golang+net+https",nil);
	if err != nil {
		log.Fatal(err);
	}

	resp, err := client.Do(req);
	if err != nil {
		log.Fatal(err);
	}

	defer resp.Body.Close();
	log.Println("Status code", resp.StatusCode);

}




