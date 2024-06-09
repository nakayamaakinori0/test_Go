
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(n int) int {
		if n <= 1 {
			return n
		} else {
			return fibonacci()(n-1) + fibonacci()(n-2)
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
