package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetXY() (float64, float64) {
	return p.x, p.y
}

// Вычисляет дистанцию
func Distance(a, b *Point) float64 {
	x1, y1 := a.GetXY()
	x2, y2 := b.GetXY()
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	p1 := NewPoint(2.5, 4.5)
	p2 := NewPoint(1.3, 5.3)

	fmt.Println(Distance(p1, p2))
}
