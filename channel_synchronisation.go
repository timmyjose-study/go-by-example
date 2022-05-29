package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker working hard")
	time.Sleep(time.Second * 3)
	fmt.Println("Worker finished working")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done

}
