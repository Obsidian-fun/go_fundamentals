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
	Method string;
	Path string;
	Handler http.HandlerFunc;
}

// Add all route entries to RouteEntry struct, 
func (rtr *Router) Route(method string, path string, handlerFunc http.HandlerFunc) {
	 e := RouteEntry {
						Method: method,
						Path: path,
						Handler: handlerFunc,
					}
	rtr.routes = append(rtr.routes, e);
}

// Match all requested routes to RouteEntry if valid,
func (re *RouteEntry)Match(r *http.Request) bool{
	if r.Method != re.Method {
		return false;
	}

	if r.URL.Path != re.Path {
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

		e.Handler.ServeHTTP(w, r);
		return
	}

	// If not matched, return 404
	http.NotFound(w, r);
}

func main() {
	r := &Router{};

	r.Route("GET","/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("It works! Place your index.html here...\n\n"));
	});

	http.ListenAndServe(":8000",r);
}

