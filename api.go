package main

import (
	"encoding/json"
	"io"
	"log"
    "errors"
	"net/http"
)

// REF: https://tutorialedge.net/golang/consuming-restful-api-with-go/
func useApi() (Latest, error) {
	var apiEndPoint string = "https://openexchangerates.org/api/latest.json?app_id="
	var appID string = goDotEnvVariable("APP_ID")
    var result Latest
	apiEndPoint = apiEndPoint + appID

	response, err := http.Get(apiEndPoint)

	if err != nil {
		return result, errors.New("Get Request Failed")
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	writeToCache(responseData, fileName)

	json.Unmarshal([]byte(responseData), &result)

	return result, nil
}
