package main

import "fmt"

func linum() func() {
	lineNo := 0
	return func() {
		lineNo++
		fmt.Println("Current line number is", lineNo)
	}
}

func main() {
	lin := linum()
	for i := 0; i < 10; i++ {
		lin()
	}

	lin = linum()
	lin()
}
