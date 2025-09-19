package main

import (
	"fmt"
)

func main() {
	shortInt := -100
	shortFloat := 2.71828
	shortString := "строка"
	shortBool := false

	fmt.Printf("Переменная 'shortInt': значение = %d, тип = %T\n", shortInt, shortInt)
	fmt.Printf("Переменная 'shortFloat': значение = %f, тип = %T\n", shortFloat, shortFloat)
	fmt.Printf("Переменная 'shortString': значение = \"%s\", тип = %T\n", shortString, shortString)
	fmt.Printf("Переменная 'shortBool': значение = %t, тип = %T\n", shortBool, shortBool)
	fmt.Println()
}
