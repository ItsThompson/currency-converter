package main

import (
	"log"
	"os"
    "fmt"
    "strconv"

	"github.com/joho/godotenv"
)

// REF: https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
// use godot package to load/read the .env file and
// return the value of the key
func getEnvVar(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getIntEnvVar(key string) int64 {

    coonvertVal, convErr := strconv.ParseInt(getEnvVar(key), 10, 64)
    if convErr != nil {
        fmt.Println(convErr)
        os.Exit(1)
    }

    return coonvertVal
}
