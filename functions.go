package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusThree(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Printf("res = %d\n", res)

	res = plusThree(1, 2, 3)
	fmt.Printf("res = %d\n", res)

}
