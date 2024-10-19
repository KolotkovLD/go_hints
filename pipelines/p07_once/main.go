// ### 7. **Once для инициализации**
// `sync.Once` используется, когда необходимо убедиться,
// что кусок кода выполнится только один раз, обычно для инициализации.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	init := func() {
		fmt.Println("Инициализация...")
	}

	for i := 0; i < 5; i++ {
		go func() {
			once.Do(init)
		}()
	}

	// Ждем завершения горутин
	fmt.Scanln()
}
