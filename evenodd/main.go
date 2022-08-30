package main

import "fmt"

func main() {
	// create the integer slice
	is := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// as index is the same as value, just use index
	// or can use `_, value` and change i to value
	for i := range is {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
