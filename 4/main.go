package main

import (
	"fmt"
)

func main() {
	a := 20
	b := 5

	fmt.Printf("Числа: a = %d, b = %d\n", a, b)
	fmt.Printf("Сложение (a + b): %d\n", a+b)
	fmt.Printf("Вычитание (a - b): %d\n", a-b)
	fmt.Printf("Умножение (a * b): %d\n", a*b)
	fmt.Printf("Деление (a / b): %d\n", a/b)
	fmt.Printf("Остаток от деления (a %% b): %d\n", a%b)
	fmt.Println()
}