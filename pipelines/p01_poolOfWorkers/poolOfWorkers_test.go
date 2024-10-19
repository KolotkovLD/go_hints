package poolOfWorkers

import (
	"testing"
)

func TestWorker(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Создаем 3 воркера
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Отправляем 5 задач
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Собираем результаты
	for r := 1; r <= 5; r++ {
		<-results
	}
}
