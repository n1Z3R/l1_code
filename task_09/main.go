package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	sl := []int{2, 10, 5, 30, 50}
	ch := make(chan int)
	go func() {
		for _, value := range sl {
			ch <- value
			time.Sleep(2 * time.Second)
		}
		wg.Done()
	}()
	chAdd := make(chan int)
	go func() {
		for {
			x := <-ch
			chAdd <- x * 2
		}
	}()
	go func() {
		for {
			fmt.Println(<-chAdd)
		}
	}()
	wg.Wait()
}
