package main

import "fmt"

func main() {
	var rub float64 = 1000
	const USDinEUR = 0.9
	const USDinRUB = 91.5
	const EURinRUB = USDinRUB / USDinEUR

	fmt.Println(rub/USDinRUB)
	fmt.Println(rub/EURinRUB)
}