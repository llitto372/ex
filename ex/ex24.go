package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimetr() float32
}
type Rectangle struct {
	Width  float32
	Height float32
}
type Circle struct {
	Radius float32
}

func (r Rectangle) Area() float32 {
	res := r.Width * r.Height
	return res
}
func (c Circle) Area() float32 {
	res := math.Pi * c.Radius * c.Radius
	return res
}
func (r Rectangle) Perimetr() float32 {
	res := 2 * (r.Width + r.Height)
	return res
}
func (c Circle) Perimetr() float32 {
	res := 2 * math.Pi * c.Radius
	return res
}
func printShapeInfo(s Shape) {
	fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", Shape.Area(s), Shape.Perimetr(s))
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	c := Circle{Radius: 3}

	printShapeInfo(r)
	// Прямоугольник - Площадь: 50, Периметр: 30

	printShapeInfo(c)
	// Круг - Площадь: 28.27, Периметр: 18.85
}
