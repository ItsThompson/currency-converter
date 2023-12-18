package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type DataInput struct {
	Value        float64
	CurrencyFrom string
	CurrencyTo   string
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

func inputWrapper(rates map[string]float64) DataInput {
	var input DataInput
    rawInput := getInput()
    formattedInput := formatInput(rawInput)

    for i := 0; i < len(formattedInput); i++ {
		if input.Value == 0 {
			val, err := strconv.ParseFloat(formattedInput[i], 64)
			if err == nil {
				input.Value = val
			}
		}
		if checkValidCurrency(rates, formattedInput[i]) {
			currentWord := formattedInput[i]
			if input.CurrencyFrom == "" {
				input.CurrencyFrom = currentWord
			} else if input.CurrencyTo == "" {
				input.CurrencyTo = currentWord
			}
		}
	}

	// If the user didn't specify value, value is 1
	if input.Value == 0 {
		input.Value = 1
	}

	if input.CurrencyFrom == "" || input.CurrencyTo == "" {
		// Missing Inputs
		return inputWrapper(rates)
	}
	return input
}
