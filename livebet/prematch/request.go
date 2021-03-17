package prematch

import (
	"encoding/json"
	"fmt"
	"hitetbet/livebet"
	"io/ioutil"
	"log"
	"net/http"
)

var apiUrl string = "https://api.hititbet1000.com:2053/api/v1/Schedule/CategoryUpcomingEventsWithMarkets"

func GetPrematchEvents() ([]byte, error) {
	var preMatchEventParameters = map[string]interface{} {
		"CategoryId": 1,
		"CompetitionFilter": "",
		"CompetitionId": nil,
		"CountryId": nil,
		"EndDate": "2021-03-24 23:59:59",
		"EventFilter": "",
		"Limit": 0,
		"Locale": "en-gb",
		"PageIndex": 1,
		"PageSize": 200,
		"StartDate": "2021-03-09 00:00:00",
		"TeamId": 0,
		"WildcardLocale": "en-gb",
	}

	prematchRequest, requestCreateError := livebet.CreateHititbetRequest(apiUrl, preMatchEventParameters)
	if requestCreateError != nil {
		log.Println(requestCreateError, 1)
		return []byte{}, nil
	}

	client := &http.Client{}
	response, callErr := client.Do(prematchRequest)

	if callErr != nil {
		log.Println(callErr, 12312)
		return []byte{}, callErr
	}
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Error ocurred when request for prematch events: %s ", response.Status)
	}

	defer response.Body.Close()
	responseBytes, decodeErr := ioutil.ReadAll(response.Body)

	var objmap map[string]json.RawMessage
	er := json.Unmarshal(responseBytes, &objmap)
	log.Println(er)
	log.Fatalln(objmap["response"])
	return responseBytes, decodeErr
}
