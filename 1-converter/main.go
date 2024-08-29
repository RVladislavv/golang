package main

import "fmt"

func main() {
	rub := getUserData()
	const USDinEUR = 0.9
	const USDinRUB = 91.5
	const EURinRUB = USDinRUB / USDinEUR

	fmt.Println("В долларах рубли",rub/USDinRUB)
	fmt.Println(rub/EURinRUB)
}

func getUserData() float64 {
	var userData float64
	fmt.Println("Введите сумму")
	fmt.Scan(&userData)
	return userData
}

func convertCurrency(sum float64, base string, target string) float64 {

}