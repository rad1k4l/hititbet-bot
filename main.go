package main

import (
	"hitetbet/livebet"
	"hitetbet/net"
)

func main() {
	go livebet.StartNotificationService()
	net.Start()
}
