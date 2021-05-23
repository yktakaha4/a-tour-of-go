package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := -1
	return func() int {
		n += 1
		return fib(n)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
