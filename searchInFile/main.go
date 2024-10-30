package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Функция для поиска фразы в файле
func searchInFile(filePath, phrase string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Не удалось открыть файл %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), phrase) {
			fmt.Printf("Найдено в файле %s, строка %d: %s\n", filePath, lineNumber, scanner.Text())
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла %s: %v\n", filePath, err)
	}
}

// Функция для обхода директории
func walkDirectory(startDir, phrase string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Используем filepath.WalkDir для обхода файлов и директорий
	err := filepath.WalkDir(startDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Ошибка при обходе %s: %v\n", path, err)
			return nil
		}

		// Если встретили директорию, запускаем новую горутину для нее
		if d.IsDir() && path != startDir {
			wg.Add(1)
			go walkDirectory(path, phrase, wg)
			return filepath.SkipDir // Пропускаем дальнейший обход этой директории в текущей горутине
		}

		// Если это файл, ищем в нем фразу
		if !d.IsDir() {
			wg.Add(1)
			go searchInFile(path, phrase, wg)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка при обходе директории %s: %v\n", startDir, err)
	}
}

func main() {
	// Получаем аргументы из командной строки
	args := os.Args[1:]

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	startDir := filepath.Dir(ex)

	//startDir := "./"
	if len(args) > 0 {
		if args[0] != "none" {
			startDir = args[0]
		}
	}
	phrase := ""
	if len(args) > 1 {
		phrase = args[1]
	} else {
		fmt.Println("Укажите фразу для поиска.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go walkDirectory(startDir, phrase, &wg)

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Поиск завершен.")
}
