package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func printArea(s Shape) {
	fmt.Printf("Area : %f \n", s.Area())
}

func main() {

	circle := Circle{Radius: 7.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	printArea(circle)
	printArea(rectangle)
}
