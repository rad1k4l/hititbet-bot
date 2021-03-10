package main

import (
	"hitetbet/livebet"
	"hitetbet/livebet/prematch"
	"hitetbet/net"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	wg.Add(1)
	go livebet.NotificationService(wg)

	wg.Add(1)
	go prematch.StartHititbetPrematchService(wg)

	wg.Add(1)
	go livebet.HubService(wg)

	wg.Add(1)
	go prematch.StartPrematchHubService(wg)

	net.StartWebsocketService()

	wg.Wait()
}
