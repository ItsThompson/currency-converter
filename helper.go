package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

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

func checkValidCurrency(rates map[string]float64, currency string) bool {
	// Checks if currency string is in array
	for key := range rates {
		if currency == key {
			return true
		}
	}
	return false
}

func castRateFromLatest(jsonResponse Latest) map[string]float64 {
	rates := make(map[string]float64)

	for key, value := range jsonResponse.Rates {
		// Cast `any` type to `float64`
		rate, ok := value.(float64)
		if !ok {
			fmt.Println("Error: Unable to convert rate to float64")
			os.Exit(1)
		}

		rates[key] = rate
	}

	return rates
}
