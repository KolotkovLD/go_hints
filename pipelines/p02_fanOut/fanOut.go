package fanOut

import (
	"fmt"
	"sync"
	"time"
)

// Этот паттерн включает создание нескольких горутин для выполнения задач параллельно (Fan-Out)
// и затем сбор результатов от этих горутин (Fan-In).
// Это полезно, когда нужно обрабатывать несколько входных данных одновременно и объединять результаты.

func workerOut(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d завершил\n", id)
}
