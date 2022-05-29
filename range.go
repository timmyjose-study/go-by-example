package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	prod := 1
	for _, num := range nums {
		prod *= num
	}
	fmt.Println("prod:", prod)

	for idx, num := range nums {
		fmt.Printf("%d: %d\n", idx, num)
	}

	kvs := map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
