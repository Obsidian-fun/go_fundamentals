/***
Implemention of HTTP router functionality in go,
Credit: Gregory Schier, https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3

***/

package main

import (
	"net/http"
	"regexp"
	"context"
)

// hold all the route entries
type Router struct{
	routes []RouteEntry;
}

type RouteEntry struct {
	Method string;
	Path *regexp.Regexp;
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

// TODO: Review this code
// Match all requested routes to RouteEntry if valid,
func (re *RouteEntry) Match(r *http.Request) bool{
	match := re.Path.FindStringSubmatch(r.URL.Path);
	if match == nil {
		return nil;
	}

	// Create  a map to store URL parameters,
	params := make(map([string]string));
	groupNames := re.Path.SubExpNames();

	for i, group := range match {
		

	}
}

// Search for routes, else 404
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, e := range rtr.routes {
		match := e.Match(r);
		if !match {
			continue
		}
		// Got the match, serving..
		e.Handler.ServeHTTP(w, r);
		return
	}
	// If not matched, return 404
	http.NotFound(w, r);
}

func URLParam(r *http.Request, name string) string{
	ctx := r.Context();
	params := ctx.Value(name).map([string]string);

	return params[name];
}


func main() {
	r := &Router{};

	r.Route("GET","/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("It works! Place your index.html here...\n\n"));
	});

	r.Route("GET",`/hello/(?P<Message>\w))`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte,"Hello " + message);
	});

	http.ListenAndServe(":8000",r);
}

