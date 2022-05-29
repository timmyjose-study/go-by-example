package main

import "fmt"

func main() {
	messages := make(chan string, 10)
	messages <- "one"
	messages <- "two"

	close(messages)

	for msg := range messages {
		fmt.Println(msg)
	}
}
