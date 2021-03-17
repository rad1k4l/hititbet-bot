package main

import (
	"hitetbet/livebet"
	"hitetbet/livebet/prematch"
	"hitetbet/net"
	"sync"
)

func main() {
	//req, _ := proxy.NewRequest(http.MethodPost, "https://google.com", map[string]string{
	//	"data": "sfds",
	//})
	//req.AddHeader("Key1", "Value tsest")
	//req.Build()
	StartServices()
}

func StartServices() {
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
