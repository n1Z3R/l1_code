package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func typeOfInterface(v interface{}) {
	switch c := v.(type) {
	case int:
		fmt.Println("int:", c)
	case string:
		fmt.Println("string:", c)
	case bool:
		fmt.Println("bool:", c)
	case chan string:
		fmt.Println("chan:", c)
	}
}

func main() {
	str := "Golang"
	typeOfInterface(str)
	n := 10
	typeOfInterface(n)
	b := false
	typeOfInterface(b)
	ch := make(chan string)
	typeOfInterface(ch)

}
