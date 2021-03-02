package livebet

import (
	"github.com/gorilla/websocket"
	"sync"
)

var clientLock *sync.Mutex

func init() {
	clientLock = &sync.Mutex{}
}

var clients []*websocket.Conn

func HubService() {
	for data := range LiveBettingCh {
		for index, client := range clients {
			go func(cl *websocket.Conn, i int) {
				writeError := cl.WriteMessage(websocket.TextMessage, data)
				if writeError != nil {
					deleteClient(i)
					cl.Close()
				}
			}(client, index)
		}
	}
}

func deleteClient(index int) {
	clientLock.Lock()
	defer clientLock.Unlock()
	clients = append(clients[:index], clients[index+1:]...)
}

func AddClient(client *websocket.Conn) {
	clientLock.Lock()
	defer clientLock.Unlock()
	_ = client.WriteMessage(websocket.TextMessage, GetActualData())
	clients = append(clients, client)
}

func ClientCount() int {
	clientLock.Lock()
	defer clientLock.Unlock()
	return len(clients)
}

