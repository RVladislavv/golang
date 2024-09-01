package main

import "fmt"

const USDinEUR = 0.9
const USDinRUB = 91.5

func main() {
	base, sum, target := getUserData()
	res := convertCurrency(sum, base, target)
	fmt.Printf("%f %s в %s: %f ", sum, base, target, res)
}

func getUserData() (string, float64, string) {
	var baseCurrency string
	var sumValue float64
	var targetCurrenct string
	for {
		fmt.Println("Введите исходную валюту: RUB | USD | EUR")
		fmt.Scan(&baseCurrency)
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

func convertCurrency(sum float64, base string, target string) float64 {
	var res float64
	if base == "USD" {
		if target == "EUR" {
			res = sum * USDinEUR
		}
		if target == "RUB" {
			res = sum * USDinRUB
		}
	}
	if base == "EUR" {
		if target == "USD" {
			res = sum / USDinEUR
		}
		if target == "RUB" {
			res = sum / USDinEUR * USDinRUB
		}
	}
	if base == "RUB" {
		if target == "USD" {
			res = sum / USDinRUB
		}
		if target == "EUR" {
			res = sum / USDinRUB * USDinEUR
		}
	}
	return res
}
