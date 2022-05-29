package main

import "fmt"

func factorial(n int) int {
	if n <= 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var num int
	fmt.Scanf("%d", &num)

	fmt.Printf("factorial(%d) = %d\n", num, factorial(num))

	var fib func(int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Printf("{fib(%d) = %d", num, fib(num))
}