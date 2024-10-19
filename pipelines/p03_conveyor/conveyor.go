package conveyor

// Паттерн конвейера используется для обработки данных через последовательность стадий.
// Каждая стадия получает входные данные, обрабатывает их и передает результат на следующую стадию.
// Этот паттерн полезен, когда данные проходят через серию преобразований или этапов.

// Стадия 1: Генерация чиселs
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Стадия 2: Умножение чисел
func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// Стадия 3: Возведение чисел в квадрат
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
