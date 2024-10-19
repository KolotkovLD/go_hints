package fanOut

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	var wg sync.WaitGroup

	// Fan-Out
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerOut(i, &wg)
	}

	// Ожидаем завершения всех воркеров
	wg.Wait()
}
