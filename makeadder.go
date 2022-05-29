package main

import "fmt"

func makeAdder(n int) func(int) int {
	return func(m int) int {
		return m + n
	}
}

func main() {
	add10 := makeAdder(10)
	fmt.Println(add10(100))

	add5 := makeAdder(5)
	fmt.Println(add5(200))
}
