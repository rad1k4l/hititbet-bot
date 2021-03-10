package prematch

import (
	"bytes"
	"crypto/md5"
	"log"
	"sync"
	"time"
)

var pollInterval time.Duration = 1 * time.Minute
var actualData []byte
var previousHash []byte
var NewEventCh chan []byte
var lock *sync.Mutex

func init() {
	lock = &sync.Mutex{}
	NewEventCh = make(chan []byte, 1)
}

// Atomic set actual data
func setActualData(actual []byte) {
	lock.Lock()
	defer lock.Unlock()
	actualData = actual
}

func GetActualData() []byte {
	lock.Lock()
	defer lock.Unlock()
	return actualData
}

func StartHititbetPrematchService(wg *sync.WaitGroup) {
	defer wg.Done()
	for true {
		resp, e := GetPrematchEvents()
		if e != nil {
			log.Println(e)
			return
		}

		actualHash := md5.Sum(resp)

		setActualData(resp)
		if res := bytes.Compare(previousHash, actualHash[:]); res != 0 {
			// data changed
			NewEventCh <- resp
			previousHash = actualHash[:]
		}

		time.Sleep(pollInterval)
	}
}
