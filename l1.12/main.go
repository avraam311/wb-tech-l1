package main

import "fmt"

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}
	mapa := make(map[string]bool)
	var res []string

	// проходимся по слайсу значений и проверяем есть ли в мапе. Если в мапе нет, значит значение уникальное и
	// мы добавляем его в итоговый слайс и в мапу в качестве ключа, чтобы в следующий раз проверять другие
	// элементы на уникальность. Если в мапе есть просто пропускаем элемент.
	for _, animal := range animals {
		_, ok := mapa[animal]
		if !ok {
			res = append(res, animal)
			mapa[animal] = true
		}
	}

	fmt.Println(res)
}
