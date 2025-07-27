// объявление главного пакета main
package main

// импорт нужных пакетов
import "fmt"

// объявление структуры Human с 3 полями и 2 методами
type Human struct {
	name   string
	age    int
	weight int
}

// объявление структуры  Action со встраиванием методов Human
type Action struct {
	Human
}

// метод структуры Human для приветствия через поле name
func (h Human) greet() {
	fmt.Printf("-Hello, my name is %v\n", h.name)
}

// метод структуры Human, который расчитывает и возвращает год рождения человека
func (h Human) birthYear() int {
	return 2025 - h.age
}

// объявление главной функции(точка входа)
func main() {
	// инициализция структуры Human
	h := Human{
		name:   "Alex",
		age:    25,
		weight: 70,
	}
	// инициализация структуры Action
	a := Action{
		Human: h,
	}

	// вызов метода greet структуры Human через переменную a типа Action, который надо доступен из-за встраивания методов
	a.greet()
	// вывод результата вызова метода birthYear структуры Human через переменную a типа Action, который доступен из-за встраивания методов
	fmt.Printf("-What year were you born in, %v ?\n-In %v\n", a.name, a.birthYear())
}
