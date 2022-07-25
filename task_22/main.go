package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

func main() {
	a := new(big.Int)
	a.SetString("1000000000000000000", 10)
	b := new(big.Int)
	b.SetString("900000000000000000", 10)
	res := new(big.Int)
	fmt.Printf("Сложение %d\n", res.Add(a, b))
	fmt.Printf("Вычитание %d\n", res.Sub(a, b))
	fmt.Printf("Умножение %d\n", res.Mul(a, b))
	fmt.Printf("Деление %d\n", res.Div(a, b))
}
