package main

import "fmt"

func main() {
	input := []int{49, 0, 1, 0, 2}

	mostWater := findContainer(input)

	fmt.Println("most water: ", mostWater)
}

func findContainer(m []int) int {
	area := 0

	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			a := getArea(i, m[i], j, m[j])
			if a > area {
				area = a
			}
		}
	}

	return area
}

func getArea(i int, v1 int, j int, v2 int) int {
	y := 0

	x := j - i

	if v2 > v1 {
		y = v1
	} else {
		y = v2
	}

	return x * y
}
