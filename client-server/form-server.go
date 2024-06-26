/***
Simple form submission handling..
***/

package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Form</title>
</head>
<body>
 <form action="#">
  <input type="text", name="in">
  <input type="submit", value="Submit">
 </form>
</body>
</html>
`

func Greetings(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>Hello World</h1>");
}

func FormServer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","text/HTML");
	switch r.Method {
		case "GET":
			io.WriteString(w,form);
		case "POST":
			r.ParseForm();
			value := r.FormValue("in");
			io.WriteString(w, value);
	}
}

func main() {
	r := http.NewServeMux();

	r.HandleFunc("/", Greetings);
	r.HandleFunc("/form", FormServer);

	err := http.ListenAndServe(":3000",r); if err != nil {
		log.Fatal(err);
	}

}

