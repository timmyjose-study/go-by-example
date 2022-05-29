package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums: ", nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("Total:", total)
}

func main() {
	sum()
	sum(1)
	sum(1, 2)
	sum(1, 2, 3)

	mySlice := []int{10, 11, 12, 13, 14, 15}
	sum(mySlice...)

}
