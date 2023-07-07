package main

import "fmt"

func main() {
	input := []int{49, 0, 1, 0, 2}

	mostWater := findContainer(input)

	fmt.Println("most water: ", mostWater)
}

func findContainer(m []int) int {
	containers := []int{}

	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			fmt.Println("checking m[", i, "]=", m[i], " and m[", j, "]=", m[j])

			a := getArea(i, m[i], j, m[j])
			containers = append(containers, a)

			fmt.Println("got area = ", a)
		}
	}

	fmt.Println("resulting containers array: ", containers)

	if len(containers) < 1 {
		return 0
	}

	return findMax(containers)
}

func getArea(i int, v1 int, j int, v2 int) int {
	x := 0
	y := 0

	x = j - i

	if v2 > v1 {
		y = v1
	} else {
		y = v2
	}

	return x * y
}

func findMax(c []int) int {
	max := c[0]

	for _, n := range c {
		if n > max {
			max = n
		}
	}

	return max
}
