package main

import (
	"fmt"
	"math"
)

// структура point с полями x и y. Точка.
type point struct {
	x float64
	y float64
}

// конструктор для создания новой точки
// возвращаем указатель для экономии памяти и возможности изменять структуру напрямую
func newPoint(x, y float64) *point {
	return &point{x, y}
}

// метод для вычисления расстояния между точками по теореме пифагора
func (p *point) Distance(other *point) float64 {
	deltaX := p.x - other.x
	deltaY := p.y - other.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	pointA := newPoint(1, 2)
	pointB := newPoint(4, 6)

	distance := pointA.Distance(pointB)
	fmt.Println(distance)
}
