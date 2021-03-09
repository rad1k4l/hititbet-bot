package prematch

import (
	"fmt"
	"hitetbet/livebet"
	"io/ioutil"
	"log"
	"net/http"
)

var apiUrl string = "https://api.hititbet1000.com:2053/api/v1/Schedule/CategoryUpcomingEventsWithMarkets"

func GetPrematchEvents() ([]byte, error) {
	var preMatchEventParameters = []byte(`{
		"CategoryId": 1,
		"CompetitionFilter": "",
		"CompetitionId": null,
		"CountryId": null,
		"EndDate": "2021-03-24 23:59:59",
		"EventFilter": "",
		"Limit": 0,
		"Locale": "en-gb",
		"PageIndex": 1,
		"PageSize": 200,
		"StartDate": "2021-03-09 00:00:00",
		"TeamId": 0,
		"WildcardLocale": "en-gb"
	}`)

	prematchRequest, requestCreateError := livebet.CreateHititbetRequest(apiUrl, preMatchEventParameters)
	if requestCreateError != nil {
		log.Println(requestCreateError)
		return []byte{}, nil
	}

	client := &http.Client{}
	response, callErr := client.Do(prematchRequest)
	if callErr != nil {
		return []byte{}, callErr
	}
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf(response.Status)
	}

	defer response.Body.Close()
	responseBytes, decodeErr := ioutil.ReadAll(response.Body)
	return responseBytes, decodeErr

}
