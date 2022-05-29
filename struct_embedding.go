package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base // embedded
	str  string
}

func main() {
	co := container{
		base: base{num: 42},
		str:  "container",
	}

	fmt.Printf("co = { num = %v, str = %v }\n", co.num, co.str) // co.num works due to embedding
	fmt.Println("Also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe()) // this works since base implements the interface

}