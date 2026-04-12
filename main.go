package main

import (
	"fmt"
	"math"
)

// Интерфейс
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Круг
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Функция для вывода информации
func printShapeInfo(s Shape) {
	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("Прямоугольник - Площадь: %.2f, Периметр: %.2f\n", v.Area(), v.Perimeter())
	case Circle:
		fmt.Printf("Круг - Площадь: %.2f, Периметр: %.2f\n", v.Area(), v.Perimeter())
	default:
		fmt.Printf("Фигура - Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimeter())
	}
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	c := Circle{Radius: 3}

	printShapeInfo(r) // Прямоугольник - Площадь: 50.00, Периметр: 30.00
	printShapeInfo(c) // Круг - Площадь: 28.27, Периметр: 18.85
}
