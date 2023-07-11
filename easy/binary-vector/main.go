package main

import "fmt"

func main() {
	var (
		n int
		d int
	)

	fmt.Scanln(&n)

	vector := make([]int, n)

	for i := 0; i < 1000; i++ {
		_, err := fmt.Scanln(&d)
		if err != nil {
			break
		}

		vector[i] = d
	}

	max := 0
	cur := 0

	for _, i := range vector {
		if i == 1 {
			cur++

			if cur > max {
				max = cur
			}
		} else {
			cur = 0
		}
	}

	fmt.Println(max)
}
