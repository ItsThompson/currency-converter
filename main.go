package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	// "net/http"
	"os"

	"github.com/joho/godotenv"
)

type Latest struct {
    Timestamp int `json:"timestamp"`
    Rates map[string]any `json:"rates"`
}

type Rates struct {
    Currency string
    Rate float64
}

// REF: https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// REF: https://tutorialedge.net/golang/consuming-restful-api-with-go/
func getLastestExchangeRates() Latest {
	/* var apiEndPoint string = "https://openexchangerates.org/api/latest.json?app_id="
	var appID string = goDotEnvVariable("APP_ID")
	apiEndPoint = apiEndPoint + appID

	response, err := http.Get(apiEndPoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	} */

    // Start of Test Code

    jsonFile, err := os.Open("example.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    responseData, _ := io.ReadAll(jsonFile)

    // End of Test Code

	var result Latest
	json.Unmarshal([]byte(responseData), &result)

	return result
}

func getRateFromApi(jsonResponse Latest) map[string]float64 {
    // map[KeyType]ValueType
    rates := make(map[string]float64)

	for key, value := range jsonResponse.Rates {
        // Each value is an `any` type, that is type asserted as a float64 
        rate, ok := value.(float64)
		if !ok {
			fmt.Println("Error: Unable to convert rate to float64")
            os.Exit(1)
		}

        rates[key] = rate
	}

    return rates
}


func main() {
	jsonResponse := getLastestExchangeRates()
    rates := getRateFromApi(jsonResponse)
    _ = rates
    timestamp := jsonResponse.Timestamp
    _ = timestamp


    var currencyFrom string
    var currencyTo string

    fmt.Print("Convert From: ")
    fmt.Scan(&currencyFrom)
    fmt.Print("Convert To: ")
    fmt.Scan(&currencyTo)
}
