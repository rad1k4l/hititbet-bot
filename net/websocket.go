package net

import (
	"flag"
	"github.com/gorilla/websocket"
	"hitetbet/livebet"
	"log"
	"net/http"
)

var addr = flag.String("addr", "0.0.0.0:8000", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	livebet.AddClient(connection)
}

func Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/socket", echo)
	//http.HandleFunc("/", home)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
