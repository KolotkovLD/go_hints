//package pipelines
//
//
//
////### 1. **Пул воркеров**
////Пул воркеров позволяет ограничить количество одновременно выполняющихся горутин, чтобы избежать перегрузки ресурсов. Задачи распределяются между воркерами для параллельной обработки.
//
//```
//
//### 2. **Fan-In и Fan-Out**
//Этот паттерн включает создание нескольких горутин для выполнения задач параллельно (Fan-Out) и затем сбор результатов от этих горутин (Fan-In). Это полезно, когда нужно обрабатывать несколько входных данных одновременно и объединять результаты.
//
//#### Пример Fan-Out:
//```go
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Printf("Воркер %d начал\n", id)
//	time.Sleep(time.Second)
//	fmt.Printf("Воркер %d завершил\n", id)
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	// Fan-Out
//	for i := 1; i <= 5; i++ {
//		wg.Add(1)
//		go worker(i, &wg)
//	}
//
//	// Ожидаем завершения всех воркеров
//	wg.Wait()
//}
//```
//
//#### Пример Fan-In:
//```go
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func worker(id int) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//		time.Sleep(time.Second)
//		out <- id * 2
//	}()
//	return out
//}
//
//func main() {
//	// Fan-In: Сбор результатов от нескольких горутин
//	result1 := worker(1)
//	result2 := worker(2)
//
//	// Вывод результатов
//	fmt.Println(<-result1)
//	fmt.Println(<-result2)
//}
//```
//
//### 3. **Конвейер (Pipeline)**
//Паттерн конвейера используется для обработки данных через последовательность стадий. Каждая стадия получает входные данные, обрабатывает их и передает результат на следующую стадию. Этот паттерн полезен, когда данные проходят через серию преобразований или этапов.
//
//```go
//package main
//
//import (
//	"fmt"
//)
//
//// Стадия 1: Генерация чисел
//func generate(nums ...int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for _, n := range nums {
//			out <- n
//		}
//		close(out)
//	}()
//	return out
//}
//
//// Стадия 2: Умножение чисел
//func multiply(in <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * 2
//		}
//		close(out)
//	}()
//	return out
//}
//
//// Стадия 3: Возведение чисел в квадрат
//func square(in <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * n
//		}
//		close(out)
//	}()
//	return out
//}
//
//func main() {
//	nums := generate(1, 2, 3, 4)
//
//	// Создание конвейера
//	multiplied := multiply(nums)
//	squared := square(multiplied)
//
//	// Потребление результатов
//	for result := range squared {
//		fmt.Println(result)
//	}
//}
//```
//
//### 4. **Select для выполнения нескольких задач одновременно**
//Оператор `select` в Go полезен, когда у вас есть несколько горутин, и вы хотите одновременно читать или записывать данные в каналы. `select` позволяет ожидать выполнения нескольких операций с каналами и продолжать работу, когда какая-либо из них готова.
//
//```go
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	c1 := make(chan string)
//	c2 := make(chan string)
//
//	go func() {
//		time.Sleep(time.Second * 1)
//		c1 <- "из c1"
//	}()
//	go func() {
//		time.Sleep(time.Second * 2)
//		c2 <- "из c2"
//	}()
//
//	for i := 0; i < 2; i++ {
//		select {
//		case msg1 := <-c1:
//			fmt.Println(msg1)
//		case msg2 := <-c2:
//			fmt.Println(msg2)
//		}
//	}
//}
//```
//
//### 5. **Паттерн Производитель-Потребитель**
//Это распространенный паттерн, в котором производитель создает данные и отправляет их потребителям. Производитель может быть горутиной, которая передает элементы в канал, а горутины-потребители забирают элементы из канала для обработки.
//
//```go
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func producer(ch chan int) {
//	for i := 0; i < 5; i++ {
//		fmt.Println("Произведено:", i)
//		ch <- i
//		time.Sleep(time.Second)
//	}
//	close(ch)
//}
//
//func consumer(ch chan int) {
//	for item := range ch {
//		fmt.Println("Потреблено:", item)
//	}
//}
//
//func main() {
//	ch := make(chan int, 5)
//	go producer(ch)
//	consumer(ch)
//}
//```
//
//### 6. **Mutex для безопасного доступа к разделяемой памяти**
//Хотя каналы в Go часто достаточны для организации конкурентности, иногда требуется управлять разделяемым состоянием напрямую. `sync.Mutex` гарантирует, что только одна горутина может одновременно обращаться к критической секции кода.
//
//```go
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	counter := 0
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			mu.Lock()
//			defer mu.Unlock()
//			counter++
//			fmt.Println("Счетчик:", counter)
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println("Итоговый счетчик:", counter)
//}
//```
//
//### 7. **Once для инициализации**
//`sync.Once` используется, когда необходимо убедиться, что кусок кода выполнится только один раз, обычно для инициализации.
//
//```go
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	var once sync.Once
//
//	init := func() {
//		fmt.Println("Инициализация...")
//	}
//
//	for i := 0; i < 5; i++ {
//		go func() {
//			once.Do(init)
//		}()
//	}
//
//	// Ждем завершения горутин
//	fmt.Scanln()
//}
//```
//
//Эти паттерны охватывают множество вариантов использования конкурентности в Go, от управления несколькими горутинами до безопасного разделения состояния и контроля выполнения задач.
