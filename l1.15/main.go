package main

import (
	"bytes"
	"fmt"
)

// функция, которая генерирует и возвращает строку, указанного размера
func createHugeString(size int) string {
	// создаем новый экземпляр буфера, чтобы эффективно генерировать строку без перераспределения памями
	// и копирования строки много раз
	buf := bytes.NewBuffer(nil)
	// проходясь по циклу добавляем новые части строки
	for i := 0; i < size; i++ {
		buf.WriteByte('x')
	}
	// собираем строку воедино и возвращаем
	return buf.String()
}

// не используем глобальную переменную
// создаем копию строки, чтобы не хранить все большую строку
func someFunc() string {
	v := createHugeString(1 << 10)
	if len(v) > 100 {
		return string([]byte(v[:100]))
	}
	return string([]byte(v))
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}

// в этом примере мы храним ссылку на огромную строку, хоть и нам нужны только первые 100 элементов этой строки,
// что несомненно приводит к утечке памяти
// var justString string

// func someFunc() {
//   v := createHugeString(1 &lt;&lt; 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }
