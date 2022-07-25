package main

import "fmt"

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//     v := createHugeString(1 << 10)
//     justString = v[:100]
// }
// func main() {
//     someFunc()
// }

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "★"
	}
	return v
}
func someFunc() {
	// При вызове v[:100] хочется получить 100 символов,
	// но по факту получаешь 100 байтов (+ можешь потерять составной байт символа),
	// поэтому нужно преобразовать строчку в слайс рун

	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])

	// Проверяем
	fmt.Println(v[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
