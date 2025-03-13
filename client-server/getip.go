/***
extracting client ip address

***/

package main

import (
	"fmt"

	"net/http"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP");
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For");
	}
	if ip == "" {
		ip = r.RemoteAddr;
	}
	fmt.Println(ip);
}

func main() {
	r := http.DefaultServeMux;
	r.HandleFunc("/", getIP);
	http.ListenAndServe(":3000",r);

}


