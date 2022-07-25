package main

import "fmt"

// Удалить i-ый элемент из слайса.

// Перемещает последний элемент слайса на i-тое место
// и возвращает отрезок слайса без последнего элемента
func delete(sl []int, i int) []int {
	sl[i] = sl[len(sl)-1]
	return sl[:len(sl)-1]
}

func main() {
	i := 1
	sl := []int{2, 4, 5, 6, 7, 9}
	sl = delete(sl, i)
	fmt.Println(sl)
}
