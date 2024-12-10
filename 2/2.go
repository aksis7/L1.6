package main

import (
	"context"
	"fmt"
	"time"
)

// worker - горутина, которая завершает работу, если контекст отменяется.
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Контекст был отменен.
			fmt.Println("Goroutine stopped due to context cancellation")
			return
		default: // Продолжаем выполнять работу.
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) 
		}
	}
}

func main() {
	// Создаем контекст с возможностью отмены.
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx) // Запускаем горутину worker с контекстом.

	time.Sleep(2 * time.Second) // Даем горутине поработать.
	cancel()                    // Отменяем контекст, сигнализируя горутине завершить работу.
	fmt.Println("Main function finished")
}
