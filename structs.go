package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 21})
	fmt.Println(person{age: 23, name: "Dave"})
	fmt.Println(newPerson("Wanda"))
	fmt.Println(&person{name: "Andy", age: 99})

	s := person{name: "Sean", age: 58}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 41
	fmt.Println(s)

}
