package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	seq := intSeq()
	for i := 0; i < 10; i++ {
		fmt.Println(seq())
	}

	seq = intSeq()
	fmt.Println(seq())
}
