package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func unique(str string) bool {
	m := make(map[string]bool)
	for _, value := range strings.ToLower(str) {
		// Если существует ключ, то возвращаем false
		if ok := m[string(value)]; ok {
			return false
		}
		// Если не существует, записываем в ключ true
		m[string(value)] = true
	}
	return true

}
func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}
