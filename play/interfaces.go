package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func PrintArea(a Shape) {
	fmt.Println(a.Area())
}

func main() {
	circle := &Circle{Radius: 4}
	square := &Square{Length: 3}
	PrintArea(circle)
	PrintArea(square)
}
