package main

import (
	"bytes"
	"fmt"
)

// функция принимает строку и возвращает ее перевернутой
func reverseString(s string) string {
	// создаем буфер для избежания утечки памяти(копирование строки в каждой итерации)
	buf := bytes.Buffer{}
	// конвертируем строку в срез рун, чтобы не терять символы unicode, которые могут занимать больше 1 байта
	runes := []rune(s)

	// проходимся по рунам с конца наперед и добавляем каждый элементв в буфер
	for i := len(runes) - 1; i >= 0; i-- {
		buf.WriteRune(runes[i])
	}

	// возвращаем готовую строку
	return buf.String()
}

func main() {
	s := "главрыба"
	reversed := reverseString(s)
	fmt.Println(reversed)
}
