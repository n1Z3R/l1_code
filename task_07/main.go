package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type SafeMap struct {
	sync.RWMutex
	content map[int]int
}

// Возвращает новый экземпляр структуры
func NewSafeMap() *SafeMap {
	return &SafeMap{content: make(map[int]int)}
}

// Безопасно записывает в мапу
func (s *SafeMap) Set(key int, val int) {
	s.Lock()
	defer s.Unlock()
	s.content[key] = val
}

// Безопасное чтение из мапы
func (s *SafeMap) Get(key int) (int, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.content[key]
	return val, ok
}

func main() {
	s := NewSafeMap()
	// Заполняет конкурентно мапу данными
	func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(i int) {
				s.Set(i, rand.Intn(100))
				wg.Done()
			}(i)
		}
		wg.Wait()
	}()
	fmt.Println(s.Get(5))
}
