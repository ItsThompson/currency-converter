package main

import "math"

func convert(dataInput DataInput, rates map[string]float64) float64 {

	var conversion float64

	var convertFromUsd = func(value float64, targetRate float64) float64 {
		var convertedValue float64 = value * targetRate
		return convertedValue
	}

	if dataInput.CurrencyFrom == "USD" {
		// Convert Base Currency to Target Currency
		conversion = convertFromUsd(dataInput.Value, rates[dataInput.CurrencyTo])
	} else {
		// Convert Origin Currency to Base Currency and then Convert to Target Currency
		rateFrom := rates[dataInput.CurrencyFrom]
		var value float64 = dataInput.Value / rateFrom
		conversion = convertFromUsd(value, rates[dataInput.CurrencyTo])
	}

	return math.Round(conversion*(math.Pow10(precision))) / math.Pow10(precision)
}
