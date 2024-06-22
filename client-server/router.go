/***
Implemention of HTTP router functionality in go,
Credit: https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3

***/

package main

import (
	"net/http"
)

type Router struct{
	routes []RouteEntry;
}

type RouteEntry struct {
	URL string;
	method string;
	Handler http.HandlerFunc;
}

// Matching end-point,
func (rtr *Router) Route(url string, method string, handlerFunc http.HandlerFunc) {
	var e := new RouteEntry {
						URL : url,
						Method: method,
						Handler: handlerFunc,
					}
	rtr.routes = append(rtr.routes, e);
}

// 404 Not Found
func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r);
}

func main() {
	r := &Router{};
	http.ListenAndServe(":8000",r);
}

