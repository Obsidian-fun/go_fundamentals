/***
Implemention of HTTP router functionality in go,
Credit: Gregory Schier, https://codesalad.dev/blog/how-to-build-a-go-router-from-scratch-3

***/

package main

import (
	"net/http"
	"log"
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
// ^ in regex indicates the start of a string and $ indicates the end
	exactPath := regexp.MustCompile("^" + path + "$");

	 e := RouteEntry {
					Method:   method,
					Path:     exactPath,
					Handler:  handlerFunc,
				}
	rtr.routes = append(rtr.routes, e);
}

// TODO: Review this code
// Match all requested routes to RouteEntry if valid,
func (re *RouteEntry) Match(r *http.Request) map[string]string{
	match := re.Path.FindStringSubmatch(r.URL.Path);
	if match == nil {
		return nil;
	}

	// Create  a map to store URL parameters,
	params := make(map[string]string);
	groupNames := re.Path.SubexpNames();

	for i, group := range match {
		params[groupNames[i]] = group;
	}
	return params;
}

// Search for routes, else 404
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", r)
			http.Error(w, "Uh oh!", http.StatusInternalServerError);
		}
	}()

	for _, e := range rtr.routes {
		params := e.Match(r);
		if params == nil{
			continue
		}
		// Create new request with params stored in context,
		ctx := context.WithValue(r.Context(),"params", params);
		e.Handler.ServeHTTP(w, r.WithContext(ctx));
		return
	}
	// If not matched, return 404
	http.NotFound(w, r);
}

func URLParam(r *http.Request, name string) string{
	ctx := r.Context();
	params := ctx.Value("params").(map[string]string);

	return params[name];
}


func main() {
	r := &Router{};

	r.Route("GET","/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("It works! Place your index.html here...\n\n"));
	});

	r.Route("GET",`/hello/(?P<Message>\w+)`, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + URLParam(r, "Message") + "\n"));
	});

	r.Route("GET","/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Something Bad happened!");
	});

	http.ListenAndServe(":8000",r);
}

