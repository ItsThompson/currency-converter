package main

import (
	"fmt"
	"time"
    "os"
)


//TODO: .env variables and place globals closer to concern
var fileName string = "cache.json"
var cacheExpiry int64 = 300
var precision int = 2


func main() {
    for true {
        var now int64 = time.Now().Unix()

        rates,err := caller(now)

        if err != nil {
            fmt.Println("Error: No currency exchange data found")
            os.Exit(1)
        }

        inputData := inputWrapper(rates)
        fmt.Println(convert(inputData, rates))
    }
}
