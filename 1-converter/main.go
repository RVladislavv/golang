package main

import (
	"fmt"
	"strings"
)

const USDinEUR = 0.9
const USDinRUB = 91.5

var convertMap = &map[string]map[string]float64 {
	"USD": {
		"EUR": USDinEUR,
		"RUB": USDinRUB,
	},
	"EUR": {
		"USD": 1 / USDinEUR,
		"RUB": USDinRUB / USDinEUR,
	},
	"RUB": {
		"USD": 1 / USDinRUB,
		"EUR": USDinEUR / USDinRUB,
	},
}

func main() {
	base, sum, target := getUserData()
	res := convertCurrency(sum, base, target, convertMap)
	fmt.Printf("%.2f %s в %s: %.2f ", sum, base, target, res)
}

func getUserData() (string, float64, string) {
	var baseCurrency string
	var sumValue float64
	var targetCurrenct string
	for {
		fmt.Println("Введите исходную валюту: RUB | USD | EUR")
		fmt.Scan(&baseCurrency)
		baseCurrency = strings.ToUpper(baseCurrency)
		if baseCurrency == "RUB" || baseCurrency == "USD" || baseCurrency == "EUR" {
			break
		} else {
			fmt.Println("Вы ввели неправильно валюту, попробуйте ещё раз")
		}
	}
	for {
		fmt.Println("Введите сумму:")
		fmt.Scan(&sumValue)
		if sumValue >= 0 {
			break
		} else {
			fmt.Println("Вы ввели неправильно сумму, попробуйте ещё раз")
		}
	}
	for {
		fmt.Println("Введите целевую валюту(отличную исходной): RUB | USD | EUR")
		fmt.Scan(&targetCurrenct)
		targetCurrenct = strings.ToUpper(targetCurrenct)

		if baseCurrency == targetCurrenct {
			fmt.Println("Целевая валюта идентична исходной, выберите другую")
			continue
		}
		if targetCurrenct == "RUB" || targetCurrenct == "USD" || targetCurrenct == "EUR" {
			break
		} else {
			fmt.Println("Вы ввели неправильно валюту, попробуйте ещё раз")
		}
	}
	return baseCurrency, sumValue, targetCurrenct
}

func convertCurrency(sum float64, base string, target string, convMap *map[string]map[string]float64) float64 {
	value := (*convMap)[base][target]
	return value * sum
}
