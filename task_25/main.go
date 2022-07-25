package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

// Ждет истечения времени и отправляет текущее время по каналу
func sleep(t time.Duration) {
	<-time.After(t)
}

// Ждет истечения времени контекста
func sleepTimeout(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	<-ctx.Done()
}

func main() {
	now := time.Now()
	sleep(2 * time.Second)
	fmt.Println("ended", time.Since(now))
	sleepTimeout(2 * time.Second)
	fmt.Println("ended", time.Since(now))
}
