package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

// automatica implementaion of geometry for circle by implementing all
// of geometry's methods

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

// similarly for the rect type

type rect struct {
	width, height int
}

func (r *rect) area() float64 {
	return float64(r.width * r.height)
}

func (r *rect) perimeter() float64 {
	return float64(2 * (r.width + r.height))
}

// this function can accept both circle and rect since they both
// "implement" the geometry interface

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

func main() {
	r := rect{width: 10, height: 20}
	measure(&r)

	c := circle{radius: 10.0}
	measure(&c)
}
