package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var res float64
	operation, userNumbers := inputData()
	fmt.Println(operation, userNumbers)
	res = calculateOperations(operation, userNumbers)
	fmt.Printf("%.2f", res)
}

func inputData() (string, string) {
	var operation string

	fmt.Println("Введите операцию(AVG - среднее, SUM - сумму, MED - медиану):")
	fmt.Scanln(&operation)

	fmt.Println("Введите неограниченное число чисел через запятую(2, 9, 10...)")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	return operation, text
}

func calculateOperations(operation, data string) float64 {
	var res float64
	var numbers []float64

	numbersStr := strings.Split(data, " ")

	for _, el := range numbersStr {
		el = strings.TrimSpace(el)
		el = strings.Trim(el, ",")
		num, _ := strconv.ParseFloat(el, 64)
		numbers = append(numbers, num)
	}

	var sum float64
	var length = len(numbers)

	for _, el := range numbers {
		sum += el
	}
	if operation == "AVG" {
		res = sum / float64(length)
	}
	if operation == "SUM" {
		res = sum
	}
	if operation == "MED" {
		res = calculateMedian(numbers)
	}

	return res
}

func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)

	length := len(numbers)

	if length == 0 {
		return 0
	}
	if length%2 == 1 {
		return numbers[length/2]
	}
	return (numbers[length/2-1] + numbers[length/2]) / 2
}
