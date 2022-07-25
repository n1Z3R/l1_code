package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func main() {
	ch := make(chan int)
	count := 0
	fmt.Scan(&count)

	// Создание контекста с таймером
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(count)*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	// Пишет в канал каждую секунду
	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)

	// Читает с канала
	go func() {
		for {
			select {
			// При окончании таймаута завершает работу
			case <-ctx.Done():
				defer wg.Done()
				fmt.Printf("Конец, прошло %v сек\n", count)
				return
			case msg := <-ch:
				fmt.Println(msg)
			}
		}
	}()

	wg.Wait()
}
