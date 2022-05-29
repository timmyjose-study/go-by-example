package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func swap(i, j *int) {
	t := *i
	*i = *j
	*j = t
}

func main() {
	i := 1
	fmt.Println("i:", i)

	zeroval(i)
	fmt.Println("After zeroval, i:", i)

	zeroptr(&i)
	fmt.Println("After zeroptr, i:", i)

	n, m := 1, 2
	fmt.Printf("Before swap, n = %d, m = %d\n", n, m)

	swap(&n, &m)
	fmt.Printf("After swap, n = %d, m = %d\n", n, m)

}