package main

import "os"

func main() {
	panic("something went terribly wrong")

	_, err := os.Create("/tmp/tempfile")
	if err != nil {
		panic("could not create temp file")
	}

}