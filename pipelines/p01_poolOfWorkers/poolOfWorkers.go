package poolOfWorkers

import (
	"fmt"
	"time"
)

// ### 1. **Пул воркеров**
// Пул воркеров позволяет ограничить количество одновременно выполняющихся горутин, чтобы избежать перегрузки ресурсов. Задачи распределяются между воркерами для параллельной обработки.

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Воркер %d начал задачу %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Воркер %d завершил задачу %d\n", id, j)
		results <- j * 2
	}
}
