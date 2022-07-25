package main

import (
	"fmt"
	"math/rand"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.>

func qsort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}

	left := 0
	right := len(sl) - 1

	// Выбирает опорный элемент
	pivot := rand.Int() % len(sl)

	// Меняет местами опорный и самый правый элемент
	sl[pivot], sl[right] = sl[right], sl[pivot]

	// Перемещает элементы меньше опорного налево
	for i := range sl {
		if sl[i] < sl[right] {
			sl[i], sl[left] = sl[left], sl[i]
			left++
		}
	}

	// Ставит опорный элемент после последнего левого
	sl[left], sl[right] = sl[right], sl[left]

	// Рекурсия левой части
	qsort(sl[:left])
	// Рекурсия правой части без опорного элемента
	qsort(sl[left+1:])

	return sl
}

func main() {
	fmt.Println(qsort([]int{5, 3, 10, 22, 8, 1, 30}))
}
