package conveyor

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	nums := generate(1, 2, 3, 4)

	// Создание конвейера
	multiplied := multiply(nums)
	squared := square(multiplied)

	// Потребление результатов
	for result := range squared {
		fmt.Println(result)
	}
}
