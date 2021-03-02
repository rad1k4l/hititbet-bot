package main

import (
	"hitetbet/livebet"
	"hitetbet/net"
)

func main() {
	go livebet.HubService()
	go livebet.NotificationService()
	net.Start()
}
