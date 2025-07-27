package main

import (
	"fmt"
	"sync"
)

// функция, которая получает массив чисел, канал для отправки и waitgroup
func generateNumbers(numbers []int, outputChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// закрываем канал после записи всех чисел
	defer close(outputChan)
	// пишем числа в канала
	for _, num := range numbers {
		outputChan <- num
	}
}

// функция, которая получает входной и выходные каналы и waitgroup
func processNumbers(inputChan <-chan int, outputChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// закрываем канал после прочтения всех чисел
	defer close(outputChan)
	// умножаем числа из входного канала на 2 и пишем их в выходной канал
	for num := range inputChan {
		outputChan <- num * 2
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	// канал для передачи чисел от генератора к обработчику
	numbersChan := make(chan int)
	// канал для передачи результатов обработки к stdout
	resultsChan := make(chan int)

	// запускаем горутину генерации чисел
	wg.Add(1)
	go generateNumbers(data, numbersChan, &wg)

	// запускаем горутину обработки чисел
	wg.Add(1)
	go processNumbers(numbersChan, resultsChan, &wg)

	// чтение из канала results и вывод в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range resultsChan {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
