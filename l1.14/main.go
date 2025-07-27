package main

import "fmt"

func defineType(el interface{}) {
	switch el := el.(type) {
	case int:
		fmt.Printf("Тип: %T\n", el)
	case string:
		fmt.Printf("Тип: %T\n", el)
	case bool:
		fmt.Printf("Тип: %T\n", el)
	case chan int:
		fmt.Printf("Тип канал: %T\n", el)
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	defineType(42)
	defineType("Hello")
	defineType(true)
	ch := make(chan int)
	defineType(ch)
	defineType(struct{}{})
}
