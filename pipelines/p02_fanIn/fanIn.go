package fanIn

import (
	"fmt"
	"time"
)

func workerIn(id int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		time.Sleep(time.Second)
		out <- id * 2
	}()
	return out
}

func mainIn() {
	// Fan-In: Сбор результатов от нескольких горутин
	result1 := workerIn(1)
	result2 := workerIn(2)

	// Вывод результатов
	fmt.Println(<-result1)
	fmt.Println(<-result2)
}
