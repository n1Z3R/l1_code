package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func BinarySearch(array []int, number int) int {
	// Переменная хранящая индекс числа после поиска
	index := -1
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (left + right) / 2

		// Если число равно середине, то возращает индекс числа
		if number == array[middle] {
			index = middle
			return index
		}

		if number < array[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	// Возвращает -1, если числа нет в упорядоченном списке
	return index
}

func main() {
	sortedSlice := []int{1, 5, 7, 12, 43, 67, 86, 92, 112}
	fmt.Println(BinarySearch(sortedSlice, 13))

}
