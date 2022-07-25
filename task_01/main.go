package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).

type Human struct {
	Name string
	Age  int
}

// Встраиваем структуру Human в Action
type Action struct {
	Human
}

// Метод human, который возращает описание экземпляра
func (h Human) About() string {
	return fmt.Sprintf("Меня зовут %s, мне %d лет.", h.Name, h.Age)
}

func main() {
	// Создаем экземпляр структуры Human и вызываем метод About
	h := Human{"Masha", 20}
	fmt.Println(h.About())

	// Создаем экземпляр структуры Action
	// и внутри инициализируем структуру Human,
	// вызываем встроенный метод About
	a := Action{Human{"Andrei", 24}}
	fmt.Println(a.About())
}
