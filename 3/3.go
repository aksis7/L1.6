package main

import (
	"fmt"
	"sync"
	"time"
)

// worker - горутина, которая завершается, если получает сигнал из канала done.
func worker(wg *sync.WaitGroup, done chan bool) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы.

	for {
		select {
		case <-done: // Получен сигнал завершения работы.
			fmt.Println("Goroutine stopped")
			return
		default: // Выполняем работу, если нет сигнала.
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) // Имитируем работу.
		}
	}
}

func main() {
	var wg sync.WaitGroup // Создаем объект WaitGroup.
	done := make(chan bool)

	wg.Add(1)            // Увеличиваем счетчик WaitGroup, так как запускаем одну горутину.
	go worker(&wg, done) // Запускаем горутину worker.

	time.Sleep(2 * time.Second) // Даем горутине поработать.
	done <- true                // Отправляем сигнал завершения.
	wg.Wait()                   // Ожидаем завершения всех горутин.
	fmt.Println("Main function finished")
}
