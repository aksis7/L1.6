package main

import (
	"fmt"
	"time"
)

// worker - горутина, которая выполняет работу до тех пор, пока не получит сигнал завершения через канал done.
func worker(done chan bool) {
	for {
		select {
		case <-done: // Если что-то приходит в канал done, горутина завершает работу.
			fmt.Println("Goroutine stopped")
			return
		default: // Если канал done пуст, продолжаем работу.
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) 
		}
	}
}

func main() {
	done := make(chan bool) // Создаем канал для передачи сигнала завершения.

	go worker(done) // Запускаем горутину worker.

	time.Sleep(2 * time.Second) // Даем горутине поработать 2 секунды.
	done <- true                // Отправляем сигнал завершения в канал.
	fmt.Println("Main function finished")
}
