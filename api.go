package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// REF: https://tutorialedge.net/golang/consuming-restful-api-with-go/
func getLastestExchangeRates() Latest {
	var apiEndPoint string = "https://openexchangerates.org/api/latest.json?app_id="
	var appID string = goDotEnvVariable("APP_ID")
	apiEndPoint = apiEndPoint + appID

	response, err := http.Get(apiEndPoint)

	if err != nil {
        // Error in Get Request -> Use Cache
        cacheData := readFromCache()
        return cacheData
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	writeToCache(responseData, fileName)

	var result Latest
	json.Unmarshal([]byte(responseData), &result)

	return result
}
