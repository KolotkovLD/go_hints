package fanOut

import (
	"fmt"
	"sync"
	"time"
)

func workerOut(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Воркер %d завершил\n", id)
}
