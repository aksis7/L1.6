package main

import (
	"fmt"
	"time"
)

// worker - горутина, которая работает до тех пор, пока не истечет таймер.
func worker(stop <-chan time.Time) {
	for {
		select {
		case <-stop: // Если наступает время из канала time.After, горутина завершает работу.
			fmt.Println("Goroutine stopped due to timeout")
			return
		default: // Выполняем работу, если таймер не сработал.
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) // Имитируем работу.
		}
	}
}

func main() {
	stop := time.After(2 * time.Second) // Создаем канал, который закроется через 2 секунды.

	go worker(stop) // Запускаем горутину worker.

	time.Sleep(3 * time.Second) // Ждем завершения горутины.
	fmt.Println("Main function finished")
}
