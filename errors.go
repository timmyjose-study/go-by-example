package main

import (
	"errors"
	"fmt"
)

func f1(n int) (int, error) {
	if n == 42 {
		return -1, errors.New("cannot work with 42")
	}

	return n + 100, nil
}

type argError struct {
	arg  int
	prob string
}

// implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s\n", e.arg, e.prob)
}

func f2(n int) (int, error) {
	if n == 42 {
		return -1, &argError{arg: n, prob: "Cannot work with 42"}
	}
	return n + 200, nil
}

func main() {
	for _, n := range []int{7, 42} {
		r, e := f1(n)
		if e != nil {
			fmt.Println("F1 failed:", e)
		} else {
			fmt.Println("f1 succeesed:", r)
		}
	}

	for _, m := range []int{7, 42} {
		r, e := f2(m)
		if e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 succeeded:", r)
		}
	}
}
