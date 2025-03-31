package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	var res float64
	res = 0.5 * r.Width * r.Height
	return res
}

func main() {
	a := Rectangle{6, 4}
	fmt.Println(a.Area())
}
