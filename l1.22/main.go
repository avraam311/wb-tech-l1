package main

import (
	"fmt"
	"math/big"
)

func main() {
	// создание огромных чисел с помощью битовых сдвигов
	a := big.NewInt((1<<20)*4 + 424)
	b := big.NewInt((1<<20)*2 + 123)

	// используем new, потому что Mul, Div, Rem, Add и Sub предназначены для работы
	// с указателями на объекты конкретного типа bit.Int
	// также new более просто и безопасно выделяет память под такие огромные числа
	// таким способом мы избегаем всяких непонятных ситуаций

	// арифметические операции с огромными числами
	product := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %v\n", product)

	quotient := new(big.Int).Div(a, b)
	remainder := new(big.Int).Rem(a, b)
	fmt.Printf("Деление: частное - %v остаток - %v\n", quotient, remainder)

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %v\n", sum)

	difference := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %v\n", difference)
}
