package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime time.Time = time.Now()
	fmt.Println("Текущая дата и время:", currentTime)
}