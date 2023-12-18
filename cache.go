package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
    "errors"
)



// NOTE: Adapts a "cache agnostic" model for calling data
//       See: https://github.com/ItsThompson/currency-converter/issues/1
func useCache(now int64, force bool) (Latest, error) {
	var latestData Latest
    fileName := getEnvVar("FILE_NAME")

	cacheExists := checkIfCacheExist(fileName)

	if cacheExists {
		cacheData := readFromCache()
		secondsElapsed := now - cacheData.Timestamp
        cacheExpiry := getIntEnvVar("CACHE_EXPIRY_IN_SECONDS")
		if force || secondsElapsed <= cacheExpiry {
			latestData = cacheData
			return latestData, nil
		} else {
			return latestData, errors.New("Cache expired")
		}
	} else {
		return latestData, errors.New("Cache does not exist")
	}
}

func checkIfCacheExist(fileName string) bool {
	_, err := os.Stat(fileName)

	// Check if the file exists
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func readFromCache() Latest {
	var latest Latest
    fileName := getEnvVar("FILE_NAME")
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	responseData, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(responseData), &latest)
	return latest
}


func writeToCache(data []byte, fileName string) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

