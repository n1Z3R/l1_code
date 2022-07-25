package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

// Преобразовывает строчку в слайс
func reverse(str string) string {
	strSlice := strings.Split(str, "")
	for i, j := 0, len(strSlice)-1; i < j; i, j = i+1, j-1 {
		// Меняет местами символы
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	return strings.Join(strSlice, "")
}

// Проходит по каждому символу и добавляет его в начало результирующей строки
func reverse2(str string) string {
	var final string
	for _, value := range str {
		final = string(value) + final
	}
	return final
}

func main() {
	fmt.Println(reverse("абырвалг"))
	fmt.Println(reverse2("абырвалг"))
}
