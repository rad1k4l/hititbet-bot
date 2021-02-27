package net

import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"hitetbet/livebet"
	"log"
	"net/http"
)

func Start() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	//handle connected
	_ = server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		log.Println("Sending welcome message")
		c.Join("chat")
		if c.IsAlive() {
			_ = c.Emit("message", livebet.GetActualData())
		}
	})
	go func() {
		for {
			server.BroadcastToAll("message", string(<-livebet.LiveBettingCh))
		}
	}()

	type Message struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)
	serveMux.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Listening on :8000")
	log.Panic(http.ListenAndServe(":8000", serveMux))

}
