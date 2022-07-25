package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//
// Реализовать все возможные способы остановки выполнения горутины.
//

// Завершается при отмене контекста
func first(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ctx.Done()
	fmt.Printf("first finished\n")
	return
}

// Завершается с помощью таймаута контекста
func second(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ctx.Done()
	fmt.Printf("second finished\n")
	return
}

//Завершается с помощью обычной записи в канал
func third(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	<-ch
	fmt.Printf("third finished\n")
	return
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Создается контекст с отменой, вызывается горутина, после вызывается отмену
	ctx, cancel := context.WithCancel(context.Background())
	go first(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()

	// Создается контекст с таймаутом
	ctxWT, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go second(ctxWT, &wg)

	// Создается канал, передается в горутину, после в канал передается пустая структура, которая завершает работу горутины
	ch := make(chan struct{})
	go third(&wg, ch)
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
	wg.Wait()

}
