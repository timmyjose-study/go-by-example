package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}

	for j := 3; j <= 10; j++ {
		fmt.Println(j)
	}

	for { // infinite loop
		fmt.Println("This is an infinite loop!")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println(n)
			continue
		}
	}
}
