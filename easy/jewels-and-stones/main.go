package main

import (
	"fmt"
)

func main() {
	var (
		j string
		s string
	)

	fmt.Scanf("%s", &j)
	fmt.Scanf("%s", &s)

	counter := 0

	jewelMap := make(map[rune]bool, 0)

	for _, jewel := range j {
		jewelMap[jewel] = true
	}

	for _, stone := range s {
		if _, ok := jewelMap[stone]; ok {
			counter++
		}
	}

	fmt.Println(counter)
}
