package main

// **Select для выполнения нескольких задач одновременно**
// Оператор `select` в Go полезен, когда у вас есть несколько горутин,
// и вы хотите одновременно читать или записывать данные в каналы.
// `select` позволяет ожидать выполнения нескольких операций с каналами
// и продолжать работу, когда какая-либо из них готова.

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "из c1"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "из c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
