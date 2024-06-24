/***
Implemention of HTTP router functionality in go,
Credit: https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3

***/

package main

import (
	"net/http"
)

// hold all the route entries
type Router struct{
	routes []RouteEntry;
}

type RouteEntry struct {
	URL string;
	Method string;
	Handler http.HandlerFunc;
}

// Add all route entries to RouteEntry struct, 
func (rtr *Router) Route(url string, method string, handlerFunc http.HandlerFunc) {
	 e := RouteEntry {
						URL : url,
						Method: method,
						Handler: handlerFunc,
					}
	rtr.routes = append(rtr.routes, e);
}

// Match all requested routes to RouteEntry if valid,
func (re *RouteEntry)Match(r *http.Request) {
	if r.Method != re.Method {
		return false;
	}

	if r.Handle != re.Handle {
		return false;
	}

	return true;
}

// Search for routes, else 404
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, e := range rtr.routes {
		match := e.Match(r);
		if !match {
			continue
		}

		e.HandlerFunc.ServeHTTP(w, r);
		return
	}

	// If not matched, return 404
	http.NotFound(w, r);
}

func main() {
	r := &Router{};
	http.ListenAndServe(":8000",r);
}

