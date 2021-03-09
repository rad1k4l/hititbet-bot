package main

import (
	"hitetbet/livebet"
	"hitetbet/livebet/prematch"
	"hitetbet/net"
	"log"
)

func main() {

	go livebet.NotificationService()
	go prematch.StartHititbetPrematchService()
	go livebet.HubService()
	go prematch.StartPrematchHubService()
	net.StartWebsocketService()

	b, _ := prematch.GetPrematchEvents()
	log.Println(string(b))
}
