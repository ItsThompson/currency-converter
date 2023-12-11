package main

import (
	"errors"
	"fmt"
	"time"
)

type Latest struct {
	Timestamp int64          `json:"timestamp"`
	Rates     map[string]any `json:"rates"`
}

type DataInput struct {
	Value        float64
	CurrencyFrom string
	CurrencyTo   string
}

var fileName string = "cache.json"
var cacheExpiry int64 = 300
var precision int = 2

func getData(now int64)map[string]float64 {
	var latestData Latest
    cacheData := readFromCache()
    // Compare timestamp between now and timestamp of cache file
    secondsElapsed := now - cacheData.Timestamp
    // If difference between timestamp is over expiry, call API else use cache data
    if secondsElapsed > cacheExpiry {
        latestData = getLastestExchangeRates() 
    } else {
        latestData = cacheData
    }
	rates := castRateFromLatest(latestData)

	return rates 
}

func getCacheTimestamp() (int64, error) {
	if checkIfCacheExist(fileName) {
		cacheData := readFromCache()
        // Compare timestamp between now and timestamp of cache file
		return cacheData.Timestamp, nil
    } else {
        return 0, errors.New("No Cache")
    }
}

func main() {
    var now int64 = time.Now().Unix()
    var rates map[string]float64
    cacheExists := checkIfCacheExist(fileName)
    if cacheExists {
        rates = getData(now)
    } else {
        var latestData Latest = getLastestExchangeRates()
        rates = castRateFromLatest(latestData)
    }
    inputData := inputWrapper(rates)
    fmt.Println(convert(inputData, rates))

    for true {
        // Because we do not want to get data every loop,
        // we need to check if cache has expired. 
        now = time.Now().Unix()
        timestamp, _ := getCacheTimestamp()
        if now - timestamp > cacheExpiry {
            rates = getData(now)
        }
        inputData := inputWrapper(rates)
        fmt.Println(convert(inputData, rates))
    }
}
