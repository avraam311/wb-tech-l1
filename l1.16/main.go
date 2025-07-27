package main

import (
	"fmt"
)

// функция принмает слайс чисел и возвращает отсортированный слайс
func quickSort(arr []int) []int {
	// определяем длину слайса
	length := len(arr)
	// базовый случай рекурсии: если длина слайса меньше или равен одному, вернуть слайс
	if length <= 1 {
		return arr
	}

	// определяем индекс опорного элемента
	pivotIndex := length / 2 
	// берем опорный элемент
	pivotValue := arr[pivotIndex]
	// объявление 2 слайсов для хранения левой и правой частей слайса
	left, right := make([]int, 0), make([]int, 0) 

	// проходимся по слайсу
	for i, value := range arr {
		// если это не опорный элемент
		if i != pivotIndex {
			// если меньше опорного элементы - в левый слайс, если больше - в правый
			if value < pivotValue {
				left = append(left, value) 
			} else {
				right = append(right, value) 
			}
		}
	}

	// рекурсивно отсортировываем левую и правую части
	// и так пока весь слайс не будет отсортирован
	leftSorted := quickSort(left)
	rightSorted := quickSort(right)

	// соединяем левую часть, опорное значение и правую часть и возвращаем отсортированный слайс
	return append(append(leftSorted, pivotValue), rightSorted...)
}

func main() {
	arr := []int{3, 6, 80, 10, 1, 2, 1, -34, 100, 12}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}
