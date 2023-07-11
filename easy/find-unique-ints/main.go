package main

import (
	"fmt"
)

func main() {
	var (
		n    int
		d    int
		prev int
	)

	// number of elements
	fmt.Scanln(&n)

	if n < 1 {
		fmt.Println(0)
		return
	}

	// first number
	fmt.Scanln(&prev)

	fmt.Println(prev)

	for i := 0; i < n-1; i++ {
		fmt.Scanln(&d)

		if d != prev {
			fmt.Println(d)
			prev = d
		}
	}
}
