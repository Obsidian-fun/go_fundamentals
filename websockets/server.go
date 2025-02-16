/***
An example websocket server. 
The websocket server simply echoes back input from the client

***/
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade from http to websocket,
	conn, err := upgrader.Upgrade(w, r, nil);
	if err != nil {
		log.Println("upgrade error : ", err);
		return
	}
	defer conn.Close();
	// event loop,
	for {
		messageType, message, err := conn.ReadMessage();
		if err != nil {
			log.Println("error reading message : ",err); 
			break	
		}
		log.Printf("recieved : %s", message);
		err = conn.WriteMessage(messageType, message);
		if err != nil {
			log.Println("error writing message : ", err);
			break
		}
	}
}

// CORS rules, 
// note: to be changed on prod
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin","127.0.0.1:3000");
		w.Header().Add("Access-Control-Allow-Credentials","true");
		w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, DELETE");
	  w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With");
		next(w, r);
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html");
}

func main() {
	log.SetFlags(1); // log date not time
	r := http.DefaultServeMux;
	fs := http.FileServer(http.Dir("./static/"));

	r.Handle("/static/", http.StripPrefix("/static/", fs));
	r.HandleFunc("/socket", CORS(socketHandler));
	r.HandleFunc("/", CORS(home));

	log.Println("Connected, listening on :3000");
	log.Fatal(http.ListenAndServe(":3000", r));

}




