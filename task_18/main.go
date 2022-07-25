package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.

type Counter struct {
	sync.Mutex
	value int
}

// Создается экземпляр счетчика
func NewCounter() *Counter {
	return &Counter{}
}

// Безопасный инкремент
func (c *Counter) Add() {
	c.Lock()
	defer c.Unlock()
	c.value++

}

func main() {
	c := NewCounter()
	wg := sync.WaitGroup{}
	// Создаются инкементирующие горутины
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Add()
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}
