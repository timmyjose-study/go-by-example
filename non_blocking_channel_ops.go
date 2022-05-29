package main

import "fmt"

func main() {
	messages := make(chan string) // unbuffered
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("No messages received")
	}

	msg := "Hello, world"

	select {
	case messages <- msg:
		fmt.Println("sent message")
	default:
		fmt.Println("could not send message")
	}

	select {
	case msg := <-messages:
		fmt.Println(msg)
	case sig := <-signals:
		fmt.Println(sig)
	default:
		fmt.Println("No activity")
	}
}
