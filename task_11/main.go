package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	// Создаются мапы, симулирующие множества
	f := map[int]bool{2: true, 10: true}
	s := map[int]bool{12: true, 10: true, 2: true}

	// Выполняется проход по каждому ключу мапы и проверяется, есть ли такой в другой мапе
	for key := range f {
		if s[key] {
			fmt.Println(key)
		}

	}
}
