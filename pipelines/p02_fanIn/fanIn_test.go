package fanIn

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	// Fan-In: Сбор результатов от нескольких горутин
	result1 := workerIn(1)
	result2 := workerIn(2)

	// Вывод результатов
	fmt.Println(<-result1)
	fmt.Println(<-result2)
}
