package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // pointer receiver type
	return r.width * r.height
}

func (r rect) perimeter() int { // value receiver type
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 20}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())
}
