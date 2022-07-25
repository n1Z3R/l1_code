package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)
	for _, value := range sl {
		m[value] = true
	}
	fmt.Println(m)
}
