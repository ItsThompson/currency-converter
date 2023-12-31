package main

import (
	"fmt"
	"time"
    "os"
)

func main() {
    if checkForForce(){
        fmt.Println("Conversion is forced to update from API")
    }

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
