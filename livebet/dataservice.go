package livebet

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"sync"
	"time"
)

var previousHash []byte
var lock *sync.Mutex
var LiveBettingCh chan []byte

func init() {
	lock = &sync.Mutex{}
	LiveBettingCh = make(chan []byte, 1)
}

var actualData []byte

func GetActualData() []byte {
	lock.Lock()
	defer lock.Unlock()
	return actualData
}

func StartNotificationService() {
	for {
		start := time.Now()
		resp, e := GetLiveBetting()
		if e != nil {
			fmt.Println(e)
			return
		}

		startHash := time.Now()
		actualHash := md5.Sum(resp)

		lock.Lock()
		actualData = resp
		lock.Unlock()

		if res := bytes.Compare(previousHash, actualHash[:]); res != 0 {
			// data changed
			LiveBettingCh <- resp
			previousHash = actualHash[:]
		}
		elapsedHash := time.Since(startHash)
		log.Printf("MD5 Hash compare took %s", elapsedHash)

		elapsed := time.Since(start)

		log.Printf("Request took %s", elapsed)
	}
}
