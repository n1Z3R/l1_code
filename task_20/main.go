package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

// Преобразовывает строчку в слайс
func reverse(str string) string {
	strSlice := strings.Split(str, " ")
	for i, j := 0, len(strSlice)-1; i < j; i, j = i+1, j-1 {
		// Меняет местами слова
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	return strings.Join(strSlice, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverse(str))
}
