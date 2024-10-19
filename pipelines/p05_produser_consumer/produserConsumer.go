package main

import (
	"fmt"
	"time"
)

// ### 5. **Паттерн Производитель-Потребитель**
// Это распространенный паттерн, в котором производитель создает
// данные и отправляет их потребителям.
// Производитель может быть горутиной, которая передает элементы в канал,
// а горутины-потребители забирают элементы из канала для обработки.

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Произведено:", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch {
		fmt.Println("Потреблено:", item)
	}
}

func main() {
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
}
