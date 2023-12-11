package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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

func writeToCache(data []byte, fileName string) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

func readFromCache() Latest {
	var latest Latest
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	responseData, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(responseData), &latest)
	return latest
}
