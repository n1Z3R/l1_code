package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

func main() {
	ch := make(chan int)
	count := 0
	fmt.Println("Количество воркеров:")
	fmt.Scan(&count)

	ctx, cancel := context.WithCancel(context.Background())

	worker(ctx, count, ch)
	go func() {
		for {
			ch <- rand.Intn(100)
			time.Sleep(1 * time.Second)
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT)
		defer signal.Stop(s)
		defer close(s)
		<-s
		fmt.Printf("Finishing workers\n")
		cancel()
		close(ch)
		wg.Done()
	}()
	wg.Wait()
}
func worker(ctx context.Context, count int, ch chan int) {
	for i := 1; i <= count; i++ {
		go func(i int) {
			for {
				select {
				case v := <-ch:
					fmt.Printf("worker:%d value:%d\n", i, v)
				case <-ctx.Done():
					fmt.Printf("worker %d finished\n", i)
					return
				}
			}
		}(i)
	}
}
