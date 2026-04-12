package main

import "fmt"

// объявление структуры
type Rectangle struct {
	Width  float64
	Height float64
}

// метод для структуры
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	fmt.Println(r.Area()) // 50
}
