package main

import (
	"fmt"
	"strings"
)

// функция возвращет true, если s уникальная и false, если нет
func stringIsUnique(s string) bool {
	// объявления мапы для хранеия рун и булева значения о наличие символа в строке
	charMap := make(map[rune]bool)
	// переводим строку в нижний регистр, чтобы сделать проверку регистронезависимой
	sLower := strings.ToLower(s)

	// проходимся циклом по срезу рун
	for _, char := range sLower {
		// если символ уже есть в мапе, возвращаем false
		if charMap[char] {
			return false
		}
		// если его нет, добавляем его в мапу
		charMap[char] = true
	}
	// возвращаем true, если цикл завершился успешно
	return true

	// также можно решить это задание через слайс, но это очень неэффективно по производительности
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aAbc"}
	for _, str := range testStrings {
		result := stringIsUnique(str)
		fmt.Println(str, result)
	}
}
