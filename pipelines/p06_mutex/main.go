// ### 6. **Mutex для безопасного доступа к разделяемой памяти**
// Хотя каналы в Go часто достаточны для организации конкурентности,
// иногда требуется управлять разделяемым состоянием напрямую.
// `sync.Mutex` гарантирует, что только одна горутина может одновременно
// обращаться к критической секции кода.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
			fmt.Println("Счетчик:", counter)
		}()
	}

	wg.Wait()
	fmt.Println("Итоговый счетчик:", counter)
}
