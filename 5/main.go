package main

import (
	"fmt"
)

func calculateSumAndDifference(a, b float64) (float64, float64) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func main() {
	num1 := 15.7
	num2 := 4.2
	
	sumResult, diffResult := calculateSumAndDifference(num1, num2)

	fmt.Printf("Вызов функции для чисел %.2f и %.2f\n", num1, num2)
	fmt.Printf("Сумма: %.2f\n", sumResult)
	fmt.Printf("Разность: %.2f\n", diffResult)
	fmt.Println()
}
