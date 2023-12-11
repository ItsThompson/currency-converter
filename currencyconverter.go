package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
)

type Latest struct {
	Timestamp int            `json:"timestamp"`
	Rates     map[string]any `json:"rates"`
}

type DataInput struct {
	Value        float64
	CurrencyFrom string
	CurrencyTo   string
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
	var apiEndPoint string = "https://openexchangerates.org/api/latest.json?app_id="
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
	}

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

func checkValidCurrency(rates map[string]float64, currency string) bool {
	// Checks if currency string is in array
	for key := range rates {
		if currency == key {
			return true
		}
	}
	return false
}

func getInput() string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	rawInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return getInput()
	}

	return rawInput
}

func formatInput(input string) []string {
	var filterAlphanumeric = func(input string) string {
		var result []rune

		for _, char := range input {
			if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
				result = append(result, char)
			}
		}

		return string(result)
	}

	var separateIntegerFromNumber = func(input string) []string {
		// Initalizes at 0
		var i int
		for i < len(input) && unicode.IsDigit(rune(input[i])) {
			i++
		}

		// If the word is completely digits or completely letters
		if i == 0 || i == len(input) {
			return []string{input}
		}

		return []string{input[:i], input[i:]}
	}

	filteredInput := filterAlphanumeric(input)
	capitalizedInput := strings.ToUpper(filteredInput)
	splitInput := strings.Fields(capitalizedInput)

	// Build final input
	var formattedInput []string
	for _, word := range splitInput {
		slice := separateIntegerFromNumber(word)
		formattedInput = append(formattedInput, slice...)
	}

	return formattedInput
}

func evaluateInput(rates map[string]float64, rawInput string) DataInput {
	var input DataInput

	formattedInput := formatInput(rawInput)

	for i := 0; i < len(formattedInput); i++ {
		if checkValidCurrency(rates, formattedInput[i]) {
			currentWord := formattedInput[i]
			prevWord := formattedInput[i-1]
			prevVal, prevErr := strconv.ParseFloat(prevWord, 64)
			if prevErr == nil && input.Value == 0.0 && input.CurrencyFrom == "" {
				input.Value = prevVal
				input.CurrencyFrom = currentWord
			} else if input.CurrencyTo == "" {
				input.CurrencyTo = currentWord
			}
		}
	}
	return input
}

func getData() (map[string]float64, int) {
	// TODO: Create cache as well
	jsonResponse := getLastestExchangeRates()
	rates := getRateFromApi(jsonResponse)
	timestamp := jsonResponse.Timestamp

	return rates, timestamp
}

func convert(dataInput DataInput, rates map[string]float64) float64 {

	var conversion float64

	var convertFromUsd = func(value float64, targetRate float64) float64 {
		var convertedValue float64 = value * targetRate
		return convertedValue
	}

	if dataInput.CurrencyFrom == "USD" {
		// Base Currency
		conversion = convertFromUsd(dataInput.Value, rates[dataInput.CurrencyTo])
	} else {
		rateFrom := rates[dataInput.CurrencyFrom]
		var value float64 = dataInput.Value / rateFrom
		conversion = convertFromUsd(value, rates[dataInput.CurrencyTo])
	}

	return conversion
}

func main() {
	rates, _ := getData()
	rawInput := getInput()
	inputData := evaluateInput(rates, rawInput)
	fmt.Println(convert(inputData, rates))
}
