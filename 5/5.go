package main

import (
	"fmt"
	"sync"
	"time"
)

var stop bool        // Глобальная переменная для управления остановкой.
var mu sync.Mutex    // Мьютекс для защиты переменной stop в многопоточном окружении.

// worker - горутина, которая проверяет значение переменной stop.
func worker() {
	for {
		mu.Lock()
		if stop { // Проверяем, установлена ли переменная stop в true.
			mu.Unlock()
			fmt.Println("Goroutine stopped")
			return
		}
		mu.Unlock()
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond) // Имитируем работу.
	}
}

func main() {
	go worker() // Запускаем горутину worker.

	time.Sleep(2 * time.Second) // Даем горутине поработать.
	mu.Lock()
	stop = true // Устанавливаем переменную stop в true.
	mu.Unlock()

	time.Sleep(1 * time.Second) // Даем время горутине завершить работу.
	fmt.Println("Main function finished")
}
