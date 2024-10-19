package fanIn

import (
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
