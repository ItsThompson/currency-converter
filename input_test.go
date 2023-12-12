package main

import (
	"testing"
)

type TestData struct {
	Inputs []string
	Output DataInput
}

func TestEvaluateInput(t *testing.T) {
    cacheData := readFromCache()
	rates := castRateFromLatest(cacheData)

	allTestData := []TestData{
		{
			Inputs: []string{
				"5 gbp to SGD", "5 GBP to SGD", "5 gbp TO sgd",
				"Convert 5 gbp to SGD", "convert 5 GBP to sgd", "CONVERT 5 gbp to sgd",
				"5 GBP in SGD", "5 gbp In Sgd", "5 GBP IN SGd",
				"exchange 5 GBP to SGD", "Exchange 5 GBP to SGD", "EXCHANGE 5 GBP to Sgd",
				"What is the Equivalent of 5 GBP in SGD?", "What is the equivalent of 5 GBP in SGD?", "What is THE EQUIVALANT of 5 GBP in SGD?",
				"5 GBP to SGD CONVERSION", "5 GBP to sgd conv", "5 gbp to SGd conversion",
				"How much is 5 gbp in SGD", "How MUHC is 5 gbp in SGD", "How MUCH is 5 gbp in SGD",
				"5gbp to sgd", "5gbp in sgd", "5gbp in SGD",
				"5 gbp to sgd hkd usd", "5gbp to sgd hkd usd", "5gbp to SGd HKD USD",
			},
			Output: DataInput{Value: 5, CurrencyFrom: "GBP", CurrencyTo: "SGD"},
		},
		{
			Inputs: []string{
				"25 gbp to SGD", "25 GBP to SGD", "25 gbp TO sgd",
				"Convert 25 gbp to SGD", "convert 25 GBP to sgd", "CONVERT 25 gbp to sgd",
				"25 GBP in SGD", "25 gbp In Sgd", "25 GBP IN SGd",
				"exchange 25 GBP to SGD", "Exchange 25 GBP to SGD", "EXCHANGE 25 GBP to Sgd",
				"What is the Equivalent of 25 GBP in SGD?", "What is the equivalent of 25 GBP in SGD?", "What is THE EQUIVALANT of 25 GBP in SGD?",
				"25 GBP to SGD CONVERSION", "25 GBP to sgd conv", "25 gbp to SGd conversion",
				"How much is 25 gbp in SGD", "How MUHC is 25 gbp in SGD", "How MUCH is 25 gbp in SGD",
				"25gbp to sgd", "25gbp in sgd", "25gbp in SGD",
				"25 gbp to sgd hkd usd", "25gbp to sgd hkd usd", "25gbp to SGd HKD USD",
			},
			Output: DataInput{Value: 25, CurrencyFrom: "GBP", CurrencyTo: "SGD"},
		},
	}

	for _, testData := range allTestData {
		for _, input := range testData.Inputs {
			evalOutput := evaluateInput(rates, input)
			if testData.Output != evalOutput {
				t.Errorf(`FAILED: evaluateInput(rates, %v) Expected %v, got %v.`, input, testData.Output, evalOutput)
			} else {
				t.Logf(`PASSED: evaluateInput(rates, %v) Returned %v`, input, testData.Output)
			}
		}
	}
}
