package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.

func changeBit(num int64, i int, b bool) int64 {
	if b {
		// Установка i-ого бита числа в 1
		return num | 1<<i
	}
	// Установка i-ого бита числа в 0
	return num &^ (1 << i)
}

func main() {
	var num int64
	fmt.Println(changeBit(num, 22, false))
	fmt.Println(changeBit(num, 62, true))
}
