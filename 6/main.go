package main

import (
	"fmt"
)

func main() {
	x1, x2, x3 := 10.0, 15.0, 23.0
	
	average := (x1 + x2 + x3) / 3.0

	fmt.Printf("Числа: %.1f, %.1f, %.1f\n", x1, x2, x3)
	fmt.Printf("Среднее значение: %.2f\n", average)
	fmt.Println()
}