package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x1     float64
	y1     float64
	x2, y2 float64
}

func main() {
	c := new(Circle)
	c.x1, c.y1, c.x2, c.y2 = 1, 2, 3, 4
	c2 := Circle{x1: 0, y1: 0, x2: 0, y2: 0}
	c3 := Circle{0, 0, 5, 5}
	log(*c, c2, c3)

	r := Rectangle{
		x1: 0,
		y1: 0,
		x2: 0,
		y2: 0,
	}

	fmt.Println(r.area())
}

func log(c ...Circle) {
	fmt.Println(c)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
