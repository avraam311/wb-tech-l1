package main

import "fmt"

// reverseWords принимает строку и возвращает строку, где порядок слов перевёрнут
func reverseWords(s string) string {
	b := []byte(s)

	// переворачиваем всю строку
	reverse(b)

	start := 0
	for i := range b {
		if b[i] == ' ' {
			reverse(b[start:i]) // слово от start до i-1
			start = i + 1
		}
	}
	// последнее слово
	reverse(b[start:])

	return string(b)
}

// reverse переворачивает срез байт на месте
func reverse(b []byte) {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
}

func main() {
	s := "snow dog sun"
	s = reverseWords(s)
	fmt.Println(s) // "sun dog snow"
}
