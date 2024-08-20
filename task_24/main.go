package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

type X float64
type Y float64

type Point struct {
	x X
	y Y
}

func NewPoint(x X, y Y) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p1 Point) Distance(p Point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-p1.x), 2) + math.Pow(float64(p.y-p1.y), 2))
}

func main() {
	p1 := NewPoint(X(1), Y(3))
	p2 := NewPoint(X(2), Y(4))

	fmt.Printf("Distance = %.2F", p1.Distance(p2))
}
