package main

import (
	"fmt"
	"time"
)

// worker - горутина, которая завершает работу, если возникает паника.
func worker() {
	defer func() {
		if r := recover(); r != nil { // Перехватываем панику.
			fmt.Println("Goroutine stopped due to panic")
		}
	}()

	for {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond) // Имитируем работу.
	}
}

func main() {
	go worker() // Запускаем горутину worker.

	time.Sleep(2 * time.Second) // Даем горутине поработать.
	panic("Stop goroutine")     // Вызываем панику.
}
