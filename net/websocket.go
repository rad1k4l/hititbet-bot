package net

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"hitetbet/livebet"
)

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "0.0.0.0:8000", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		err = c.WriteMessage(1, livebet.GetActualData())
		err = c.WriteMessage(1, <-livebet.LiveBettingCh)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/socket", echo)
	//http.HandleFunc("/", home)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func _Start() {
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	http.Handle("/sock", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}

		go func() {
			defer conn.Close()
			for {
				err = wsutil.WriteServerMessage(conn, ws.OpClose, <-livebet.LiveBettingCh)
				if err != nil {
				}
			}
		}()
	}))
	_ = http.ListenAndServe(":8000", nil)
}
