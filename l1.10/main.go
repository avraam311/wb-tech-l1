package main

import (
	"fmt"
)

// функция, которая принимает массив из 8 температут и возвращает мапу с отсортированными группами
func groupTemperatures(temperatures [8]float64) map[int][]float64 {
	// инициализируем мапу
	groups := make(map[int][]float64)
	// проходимся по температурам
	for _, temp := range temperatures {
		// формируем ключи и добавляем в мапу
		key := int(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func main() {
	temps := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := groupTemperatures(temps)

	fmt.Println(result)
}
