package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Printf("%d, %d\n", a, b)

	_, c := vals()
	fmt.Printf("%d\n", c)

	d, _ := vals()
	fmt.Printf("%d\n", d)
}
