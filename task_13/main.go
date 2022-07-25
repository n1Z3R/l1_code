package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	f := 1
	s := 2
	s, f = f, s
	fmt.Println(f, s)
}
