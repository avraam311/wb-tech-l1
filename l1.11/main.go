package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	var res []int

	// берем первый элемент первого слайса
	for _, num := range a {
		// сравниваем его с каждым элементов второго слайса, и если есть совпадение, дообавляем его в
		// итоговый слайс
		for _, num2 := range b{
			if num == num2 {
				res = append(res, num)
			}
		}
	}

	fmt.Println(res)
}
