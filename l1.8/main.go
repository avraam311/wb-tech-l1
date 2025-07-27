package main

import "fmt"

// функция, которая устанавливает i-й бит числа n в 1 и возвращает n
func setOne(n int64, i int) int64 {
	n = n | 1 << i
	return n
}

// функция, которая устанавливает i-й бит числа n в 0 и возвращает n
func setZero(n int64, i int) int64 {
	n = n &^ 1 << i
	return n
}

func main() {
	// переменная типа int64 для хранения нужного числа
	var n int64
	// переменная типа int для хранения индекста нужного бита
	var i int
	fmt.Scan(&n)
	fmt.Scan(&i)
	// уменьшаем i на 1, так как в двоичном представлении биты считаются с нуля
	// чтобы избежать путаницы
	i -= 1

	// используем функцию и выводим число
	n = setZero(n, i)
	fmt.Println(n)

	// используем функцию и выводим число
	n = setOne(n, i)
	fmt.Println(n)
}
